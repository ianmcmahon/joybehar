package dcs

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/ianmcmahon/joybehar/alert"
	"golang.org/x/net/ipv4"
)

const (
	outputStreamAddr = "239.255.50.10:5010"
	inputStreamAddr  = "localhost:7778"
	maxDatagramSize  = 1024 * 1024
)

type DCSOutput interface {
	Address() uint16
	PerformAction(*dcsAgent)
}

type IntegerOutput struct {
	Addr   uint16
	Mask   uint16
	Value  uint16
	Action func(addr, value uint16)
}

func (o *IntegerOutput) Address() uint16 {
	return o.Addr
}

func (o *IntegerOutput) PerformAction(agent *dcsAgent) {
	val := toWord(agent.memory[o.Addr:o.Addr+2]) & o.Mask
	mask := o.Mask
	for mask&1 == 0 {
		mask >>= 1
		val >>= 1
	}
	if val != o.Value {
		o.Value = val
		o.Action(o.Addr, val)
	}
}

type StringOutput struct {
	Addr      uint16
	MaxLength uint16
	Value     string
	Action    func(addr uint16, value string)
}

func (o *StringOutput) Address() uint16 {
	return o.Addr
}

func (o *StringOutput) PerformAction(agent *dcsAgent) {
	val := string(agent.memory[o.Addr : o.Addr+o.MaxLength])
	if nulpos := strings.Index(val, "\x00"); val != o.Value && nulpos >= 0 {
		o.Value = val
		o.Action(o.Addr, val[0:nulpos])
	}
}

type dcsAgent struct {
	memory        []byte
	memMutex      *sync.RWMutex
	notifiers     map[uint16]DCSOutput
	notifierMutex *sync.RWMutex
	send          chan DCSMsg
}

func DCSAgent() *dcsAgent {
	agent := &dcsAgent{
		memory:        make([]byte, 0x10000),
		memMutex:      &sync.RWMutex{},
		notifiers:     make(map[uint16]DCSOutput, 0),
		notifierMutex: &sync.RWMutex{},
		send:          make(chan DCSMsg, 0),
	}
	return agent
}

func (a *dcsAgent) Register(output DCSOutput) {
	a.notifierMutex.Lock()
	if _, ok := a.notifiers[output.Address()]; !ok {
		a.notifiers[output.Address()] = output
	} else {
		fmt.Printf("Error, adding a second output notifier to the same address %p\n", output.Address())
	}
	a.notifierMutex.Unlock()
}

func (a *dcsAgent) notify(addrs []uint16) {
	a.notifierMutex.RLock()
	defer a.notifierMutex.RUnlock()

	for _, addr := range addrs {
		if notifier, ok := a.notifiers[addr]; ok {
			notifier.PerformAction(a)
		}
	}
}

type DCSMsg struct {
	Message string
	Value   string
}

func DCSMsgFromString(s string) (DCSMsg, error) {
	tokens := strings.Split(s, " ")
	if len(tokens) != 2 {
		return DCSMsg{s, ""}, fmt.Errorf("Invalid dcsMsg: %s", s)
	}
	return DCSMsg{tokens[0], tokens[1]}, nil
}

type Agent interface {
	Send(msg, val string)
	SendMsg(m DCSMsg)
	Register(output DCSOutput)
	Receive()
}

func (a *dcsAgent) Send(msg, val string) {
	a.SendMsg(DCSMsg{msg, val})
}

func (a *dcsAgent) SendMsg(m DCSMsg) {
	a.send <- m
}

func (a *dcsAgent) Receive() {
	outputUDP, err := net.ResolveUDPAddr("udp", outputStreamAddr)
	if err != nil {
		fmt.Printf("Error resolving udp addr: %v\n", err)
		panic(err)
	}
	inputUDP, err := net.ResolveUDPAddr("udp", inputStreamAddr)
	if err != nil {
		fmt.Printf("Error resolving udp addr: %v\n", err)
		panic(err)
	}

	conn, err := net.ListenUDP("udp", outputUDP)
	if err != nil {
		fmt.Printf("Error listening to udp: %v\n", err)
		panic(err)
	}

	pc := ipv4.NewPacketConn(conn)

	iface, err := multicastInterface()
	if err != nil {
		alert.Sayf("can't find specified interface %v", err)
		alert.Sayf("disabling dcsbios receive")
		return
	}
	if err := pc.JoinGroup(iface, outputUDP); err != nil {
		alert.Sayf("joining multicast group %v: %v\n", outputUDP, err)
		alert.Sayf("disabling dcsbios receive\n")
		return
	}

	if loop, err := pc.MulticastLoopback(); err == nil {
		alert.Sayf("MulticastLoobpack Status: :%v\n", loop)
		if !loop {
			if err := pc.SetMulticastLoopback(true); err != nil {
				alert.Sayf("SetMulticastLoopback error: %v\n", err)
			}
		}
	}

	go func() {
		for msg := range a.send {
			data := []byte(fmt.Sprintf("%s %s\n", msg.Message, msg.Value))
			if _, err := conn.WriteTo(data, inputUDP); err != nil {
				alert.Sayf("err in udp write: %v\n", err)
			} else {
				//fmt.Printf("sent %d bytes:%s\n", n, data)
			}
		}
	}()

	conn.SetReadBuffer(maxDatagramSize)

	for {
		buffer := make([]byte, maxDatagramSize)

		len, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			alert.Sayf("error in udp read: %v\n", err)
		}
		updatedAddrs := a.decodeDatagram(buffer[0:len])
		go a.notify(updatedAddrs)
	}
}

func isSync(b []byte) bool {
	return b[0] == 0x55 && b[1] == 0x55 && b[2] == 0x55 && b[3] == 0x55
}

func toWord(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1])<<8
}

func (a *dcsAgent) decodeDatagram(b []byte) []uint16 {
	// find sync pulse; this may not be required since all datagrams seem to start with a sync pulse
	if !isSync(b) {
		fmt.Printf("bad datagram:\n%s\n", hex.Dump(b))
	}
	b = b[4:len(b)]

	addrSeen := map[uint16]bool{}

	for {
		addr := toWord(b[0:2])
		n := toWord(b[2:4])
		if len(b) < 4+int(n) {
			fmt.Printf("%x wants %d bytes but we only have %d\n", addr, n, len(b))
			break
		}
		data := b[4 : 4+n]
		b = b[4+n : len(b)]
		if len(b) == 0 {
			break
		}
		if len(b) < 4 {
			fmt.Printf("bad datagram? all we have left is %x\n", b)
			break
		}
		//fmt.Printf("0x%.4x: %d bytes: %#v (%s)\n", addr, n, data, data)

		addrSeen[addr] = true

		a.memMutex.Lock()
		for i := uint16(0); i < n; i++ {
			a.memory[addr+i] = data[i]
		}
		a.memMutex.Unlock()
	}
	//fmt.Println()

	out := []uint16{}
	for k, _ := range addrSeen {
		out = append(out, k)
	}
	return out
}
