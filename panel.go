package main

import (
	"bufio"
	"time"

	"github.com/ianmcmahon/joybehar/alert"
	"github.com/ianmcmahon/joybehar/controls"
	"github.com/ianmcmahon/joybehar/dcs"
	"github.com/jacobsa/go-serial/serial"
)

type panelAgent struct {
	options  serial.OpenOptions
	dcsAgent dcs.Agent
	group    controls.Interceptor
}

func PanelAgent(port string, dcsAgent dcs.Agent, group controls.Interceptor) *panelAgent {
	agent := &panelAgent{
		options: serial.OpenOptions{
			PortName:        port,
			BaudRate:        250000,
			DataBits:        8,
			StopBits:        1,
			MinimumReadSize: 4,
		},
		dcsAgent: dcsAgent,
		group:    group,
	}

	return agent
}

func (a *panelAgent) Receive() {
	for {
		time.Sleep(5000 * time.Millisecond)
		port, err := serial.Open(a.options)
		if err != nil {
			continue
		}
		defer port.Close()
		alert.Say("Opening panel")

		scanner := bufio.NewScanner(port)

		for scanner.Scan() {
			msg, err := dcs.DCSMsgFromString(scanner.Text())
			if err != nil {
				alert.Sayf("%v", err)
				continue
			}
			msg, err = a.group.Intercept(msg)
			if err != nil {
				alert.Sayf("%v", err)
				continue
			}

			//fmt.Printf("%v\n", msg)
			a.dcsAgent.SendMsg(msg)
		}
		alert.Say("Panel exited")
	}
}
