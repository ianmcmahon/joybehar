package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ianmcmahon/joybehar/alert"
	"github.com/ianmcmahon/joybehar/controls"
	"github.com/ianmcmahon/joybehar/dcs"
)

var dcsAgent dcs.Agent

func main() {
	dcsAgent = dcs.DCSAgent()

	alert.Say("starting joybehar")

	warthog := controls.WarthogGroup()

	panel := PanelAgent("COM3", dcsAgent, warthog)

	// before any module is selected, we need some controls
	defaultMap := warthog.ModuleMap("")
	defaultMap.ModeToggle("stick", "paddle", controls.MODE_SHIFT, controls.MODE_NORM)
	defaultMap.Control("stick", "cms_dp").Action(controls.MODE_NORM, tempo(keyPulse(K_ESC), keyPulse(K_PAUSE)))
	defaultMap.Control("stick", "cms_dp").Action(controls.MODE_SHIFT, tempo(keyPulse(K_ESC), keyPulse(K_F10)))
	defaultMap.Control("throttle", "slewpress").Action(controls.MODE_NORM, mouseToggle(LEFT))
	defaultMap.Control("throttle", "slewpress").Action(controls.MODE_SHIFT, mouseToggle(RIGHT))
	defaultMap.Control("throttle", "tdc_aft").Action(controls.MODE_NORM, keyPulse(K_F1))
	defaultMap.Control("throttle", "tdc_left").Action(controls.MODE_NORM, keyPulse(K_F2))
	defaultMap.Control("throttle", "tdc_fwd").Action(controls.MODE_SHIFT, mouseScroll(-600))
	defaultMap.Control("throttle", "tdc_aft").Action(controls.MODE_SHIFT, mouseScroll(600))
	defaultMap.Control("throttle", "mic_up").Action(controls.MODE_NORM, keyPulse(K_LCONTROL, K_LALT, K_GRAVE))
	defaultMap.Control("throttle", "mic_dn").Action(controls.MODE_NORM, keyPress(K_GRAVE))
	defaultMap.Control("stick", "dms_up").Action(controls.MODE_NORM, keyAction(K_KPPLUS)) // discord PTT

	F5EMap(warthog)
	F14BMap(warthog)
	FA18CMap(warthog)
	F16CMap(warthog)
	JF17Map(warthog)
	ChristenEagleMap(warthog)

	dcsAgent.Register(&dcs.StringOutput{Addr: 0x0000, MaxLength: 24, Action: func(_ uint16, module string) {
		if warthog.HasModule(module) {
			alert.Sayf("mapping module %s", module)
			warthog.SetModule(module)
		} else {
			alert.Sayf("unknonn module: %s", module)
		}
	}})
	go dcsAgent.Receive()

	go panel.Receive()

	for {
		time.Sleep(1 * time.Second)
	}
}

func makeIncremental(newLabel string) controls.InterceptAction {
	return func(in dcs.DCSMsg) (dcs.DCSMsg, error) {
		val, err := strconv.ParseInt(in.Value, 10, 32)
		if err != nil {
			return in, fmt.Errorf("Error converting hsi counts to incremental: %v - %v", in, err)
		}
		if val > 0 {
			in.Value = "INC"
			in.Message = newLabel
			return in, nil
		}
		if val < 0 {
			in.Value = "DEC"
			in.Message = newLabel
			return in, nil
		}
		return in, fmt.Errorf("Zero counts don't increment or decrement")
	}
}

func mapToVals(newLabel string, vals ...string) controls.InterceptAction {
	return func(in dcs.DCSMsg) (dcs.DCSMsg, error) {
		val, err := strconv.ParseInt(in.Value, 10, 32)
		if err != nil {
			return in, fmt.Errorf("Error mapping vals, expected numeric, got %v", in)
		}
		if int(val) >= len(vals) {
			return in, fmt.Errorf("Val index bigger than the vals we can map to: %v", in)
		}
		in.Value = vals[val]
		in.Message = newLabel
		return in, nil
	}
}
