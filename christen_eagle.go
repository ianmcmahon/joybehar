package main

import "github.com/ianmcmahon/joybehar/controls"

func ChristenEagleMap(group *controls.DeviceGroup) {
	m := group.ModuleMap("Christen Eagle II")

	m.ModeToggle("stick", "paddle", controls.MODE_SHIFT, controls.MODE_NORM)
	//m.Control("stick", "trigger1").Action(controls.MODE_ALL, keyAction(K_T))
	//m.Control("stick", "trigger2").Action(controls.MODE_ALL, keyAction(K_SPACE))

	//m.Control("stick", "weaponrelease").Action(controls.MODE_NORM, keyAction(K_RALT, K_SPACE))
	//m.Control("stick", "weaponrelease").Action(controls.MODE_SHIFT, dcsToggle("RWR_PWR"))

	m.Control("stick", "nws").Action(controls.MODE_ALL, keyAction(K_S))

	//m.Control("stick", "index").Action(controls.MODE_NORM, keyPulse(K_D))
	//m.Control("stick", "index").Action(controls.MODE_SHIFT, keyPulse(K_R))

	/*
		m.Control("stick", "trim_nu").Action(controls.MODE_NORM, keyPulse(K_RCONTROL, K_PERIOD))
		m.Control("stick", "trim_nd").Action(controls.MODE_NORM, keyPulse(K_RCONTROL, K_SEMICOLON))
		m.Control("stick", "trim_lt").Action(controls.MODE_NORM, keyPulse(K_RCONTROL, K_COMMA))
		m.Control("stick", "trim_rt").Action(controls.MODE_NORM, keyPulse(K_RCONTROL, K_SLASH))

		m.Control("stick", "tms_up").Action(controls.MODE_NORM, keyPulse(K_LSHIFT, K_W))
		m.Control("stick", "tms_dn").Action(controls.MODE_NORM, keyPulse(K_LSHIFT, K_X))
		m.Control("stick", "tms_lt").Action(controls.MODE_NORM, keyPulse(K_LSHIFT, K_S))
		m.Control("stick", "tms_rt").Action(controls.MODE_NORM, keyPulse(K_LSHIFT, K_D))

		m.Control("stick", "dms_up").Action(controls.MODE_NORM, keyAction(K_KPPLUS)) // discord PTT

		m.Control("stick", "cms_up").Action(controls.MODE_NORM, keyPulse(K_RALT, K_SEMICOLON))
		m.Control("stick", "cms_dn").Action(controls.MODE_NORM, keyPulse(K_RALT, K_PERIOD))
		m.Control("stick", "cms_lt").Action(controls.MODE_NORM, keyPulse(K_RALT, K_COMMA))
		m.Control("stick", "cms_rt").Action(controls.MODE_NORM, keyPulse(K_RALT, K_MINUS))
	*/
	m.Control("stick", "cms_dp").Action(controls.MODE_NORM, tempo(keyPulse(K_R), keyPulse(K_PAUSE)))
	m.Control("stick", "cms_dp").Action(controls.MODE_SHIFT, tempo(keyPulse(K_ESC), keyPulse(K_F10)))

	m.Control("throttle", "slewpress").Action(controls.MODE_NORM, mouseToggle(LEFT))
	m.Control("throttle", "slewpress").Action(controls.MODE_SHIFT, mouseToggle(RIGHT))

	//m.Control("throttle", "mic_up").Action(controls.MODE_NORM, keyAction(K_SEMICOLON))
	//m.Control("throttle", "mic_dn").Action(controls.MODE_NORM, keyAction(K_PERIOD))
	//m.Control("throttle", "mic_fwd").Action(controls.MODE_NORM, keyAction(K_SLASH))
	//m.Control("throttle", "mic_aft").Action(controls.MODE_NORM, keyAction(K_COMMA))
	//m.Control("throttle", "mic_dp").Action(controls.MODE_NORM, keyPulse(K_ENTER))

	m.Control("throttle", "mic_up").Action(controls.MODE_SHIFT, keyPulse(K_LCONTROL, K_LALT, K_GRAVE))
	m.Control("throttle", "mic_dn").Action(controls.MODE_SHIFT, keyPress(K_GRAVE))
	m.Control("throttle", "mic_aft").Action(controls.MODE_SHIFT, keyPulse(K_BACKSLASH))

	//m.Control("throttle", "mic_dp").Action(controls.MODE_SHIFT, keyPulse(K_RSHIFT, K_K))

	//m.Control("throttle", "tdc_fwd").Action(controls.MODE_NORM, keyPulse(K_P))
	m.Control("throttle", "tdc_aft").Action(controls.MODE_NORM, keyPulse(K_F1))
	m.Control("throttle", "tdc_left").Action(controls.MODE_NORM, keyPulse(K_F2))
	m.Control("throttle", "tdc_right").Action(controls.MODE_NORM, keyPulse(K_F6))

	m.Control("throttle", "tdc_fwd").Action(controls.MODE_SHIFT, mouseScroll(-400))
	m.Control("throttle", "tdc_aft").Action(controls.MODE_SHIFT, mouseScroll(400))
	m.Control("throttle", "tdc_left").Action(controls.MODE_SHIFT, keyPulse(K_LEFTBRACE))
	m.Control("throttle", "tdc_right").Action(controls.MODE_SHIFT, keyPulse(K_RIGHTBRACE))

	//m.Control("throttle", "speedbrake").Action(controls.MODE_ALL, dcsAction("SPEED"))
	//m.Control("throttle", "boatswitch").Action(controls.MODE_ALL, dcsAction("A_FLAPS"))
	//m.Control("throttle", "chinahat").Action(controls.MODE_NORM, dcsAction("RADAR_RANGE").withVals("DEC", "", "INC"))

	//m.Control("throttle", "pinkybutton").Action(controls.MODE_NORM, dcsAction("MISSILE_UNCAGE"))
	m.Control("throttle", "pinkybutton").Action(controls.MODE_SHIFT, keyPulse(K_LCONTROL, K_E))
	//m.Control("throttle", "pinkybutton").Action(controls.MODE_SHIFT, keyPulse(K_LCONTROL, K_E))

	//	m.Control("throttle", "leftidle").Action(controls.MODE_ALL, keyToggle(keyPulse(K_RALT, K_HOME), keyPulse(K_RALT, K_END)))
	//m.Control("throttle", "rightidle").Action(controls.MODE_ALL, keyToggle(keyPulse(K_RSHIFT, K_END), keyPulse(K_RSHIFT, K_HOME)))

	//m.Control("throttle", "eacarm").Action(controls.MODE_NORM, dcsAction("FLIR_SW").withVals("1", "2"))
	//m.Control("throttle", "radalt").Action(controls.MODE_NORM, dcsAction("LST_NFLR_SW"))
	//m.Control("throttle", "apselect").Action(controls.MODE_NORM, dcsAction("LTD_R_SW"))
	//m.Control("throttle", "gearhornsilence").Action(controls.MODE_NORM, keyPulse(K_I))
	//m.Control("throttle", "gearhornsilence").Action(controls.MODE_SHIFT, keyPulse(K_U))

	m.Control("throttle", "apengage").Action(controls.MODE_NORM, dcsToggle("RWR_POWER_BTN"))
	//m.Control("throttle", "apengage").Action(controls.MODE_SHIFT, keyPulse(K_LCONTROL, K_E))

	//m.ReLabel("LG_LEVER_SWITCH", controls.MODE_NORM, "GEAR_LEVER")
	//m.ReLabel("JETT_B", controls.MODE_NORM, "SEL_JETT_BTN")
	//m.Intercept("JETT_SW", controls.MODE_NORM, mapToVals("SEL_JETT_KNOB", "1", "3", "4"))
	//m.ReLabel("PITOT_HEATER", controls.MODE_NORM, "HOOK_BYPASS_SW")
	//m.ReLabel("INLET_HEATER", controls.MODE_NORM, "ANTI_SKID_SW")
	//m.ReLabel("PITCH_DAMPER", controls.MODE_NORM, "LDG_TAXI_SW")
	//m.ReLabel("YAW_DAMPER", controls.MODE_NORM, "LAUNCH_BAR_SW")

	// knobs: FLD FLT NAV
	//  f5        ENG FORM
	//        ARM CON

	// knobs: FORM  CON INST FLD
	//  f18    POS   WARN CHRT

	//m.ReLabel("FLOOD_LIGHTS", controls.MODE_NORM, "CONSOLES_DIMMER")
	//m.ReLabel("FLIGHT_LIGHTS", controls.MODE_NORM, "INST_PNL_DIMMER")
	//m.ReLabel("NAV_LIGHTS", controls.MODE_NORM, "FLOOD_DIMMER")
	//m.ReLabel("ENGINE_LIGHTS", controls.MODE_NORM, "WARN_CAUTION_DIMMER")
	//m.ReLabel("FORMATION_LIGHTS", controls.MODE_NORM, "CHART_DIMMER")
	//m.ReLabel("CONSOLE_LIGHTS", controls.MODE_NORM, "FORMATION_DIMMER")
	//m.ReLabel("ARM_LIGHTS", controls.MODE_NORM, "POSITION_DIMMER")
	//m.ReLabel("BRI_DIM_SW", controls.MODE_NORM, "COCKKPIT_LIGHT_MODE_SW")

	//m.ReLabel("CHAFF_MODE_SEL", controls.MODE_NORM, "CMSD_DISPENSE_SW")
	//m.Intercept("FLARE_MODE_SEL", controls.MODE_NORM, mapToVals("ECM_MODE_SW", "0", "3", "4"))
}
