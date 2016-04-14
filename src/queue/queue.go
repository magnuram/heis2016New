package queue

import (
	"../config"
	"drivers"
)

func QueueOrderExists() bool {
	for i := 0; i < N_FLOORS; i++ {
		if Elinfo.ReqLocal[i] != 0 || Elinfo.ReqLocal[i] != 0 || Elinfo.ReqLocal[i] != 0 {
			return true
		}
	}
	return false
}

func QueueSetLights() {
	for i := 0; i < config.N_FLOORS; i++ {
		drivers.ElevSetButtonLamp(BUTTON_COMMAND, i, Elinfo.OrderLocal[i])
		if i != 0 {
			drivers.ElevSetButtonLamp(BUTTON_CALL_DOWN, i, Elinfo.ReqDown[i])
		}
		if i != 3 {
			drivers.ElevSetButtonLamp(BUTTON_CALL_UP, i, Elinfo.ReqUp[i])
		}
	}
}
