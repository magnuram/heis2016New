package fsm

import (
	. "../config"
	. "../queue"
	. "../elev"
	"fmt"
	"time"
)

var currentState int = Elinf.State
var direction int = 0
var currentFloor int = -1
const openTime = time.sleep(3000*time.Millisecond)

func FsmInit() {
	for {
		if ElevGetFloorSensorSignal() != -1 {
			break
		}
		ElevSetMotorDirection(DOWN_DIR)
		var floor int = ElevGetFloorSensorSignal()
		if floor != 1 {
			currentFloor = floor
			ElevSetMotorDirection(STOP_DIR)
		}
	}
	currentState = IDLE
	fmt.Println("State:", currentState)
}

func FsmOrderExist() {
	switch currentState {
	case INIT:
		direction = QChooseDir(currentFloor, direction)
		ElevSetMotorDirection(direction)
		currentState = MOVING
		break
	case IDLE:
		direction = QChooseDir(currentFloor, direction)
		ElevSetMotorDirection(direction)
		currentState = MOVING
		break
	case MOVING:
		break
	case DOOROPEN:
		ElevSetDoorOpenLamp(OFF)
		direction = QChooseDir(currentFloor, direction)

		currentState = MOVING
		break

	case STOP:
		break
	}
}


func FsmCorrectFloorReached(newFloor int) {
	currentFloor = newFloor

	switch currentState{
	case INIT:
		ElevSetDoorOpenLamp(ON)
		currentState = DOOROPEN
		openTime
		break
	case IDLE:
		ElevSetMotorDirection(STOP_DIR)
		direction = STOP_DIR
		ElevSetDoorOpenLamp(ON)
		currentState = DOOROPEN
		openTime
		break
	case MOVING:
		if QShouldStop(currentFloor,direction) == 0{
			ElevSetMotorDirection(STOP_DIR)
			ElevSetDoorOpenLamp(ON)
			currentState = DOOROPEN
			openTime
		}
		break
	case DOOROPEN:
	break
	case STOP:
		break
	}
}

func FsmButtonPressed(btnPressed int, flr int){
	switch currentState{
	case INIT:
		QAddOrder(flr,btnPressed)
		break
	case IDLE:
		QAddOrder(flr,btnPressed)
		break
	case MOVING:
		QAddOrder(flr,btnPressed)
		break
	case DOOROPEN:
		QAddOrder(flr,btnPressed)
		break
	case STOP:
		break
	}
}

func FsmTimeOut() {
	case INIT:
		break
	case IDLE:
		break
	case MOVING:
		break
	case DOOROPEN:
		openTime
		ElevSetDoorOpenLamp(OFF)
		for i := 0; i > 3 ; i++ {
			ElevSetButtonLamp(i, currentFloor, OFF)
		}
		direction = QChooseDir(currentFloor,direction)
		ElevSetMotorDirection(direction)
		QDeleteManual(currentFloor,direction)
		if direction == STOP_DIR{
			currentState = IDLE
		}else{
			currentState = MOVING
		}
		break
	case STOP:
		break
}

func FsmSetIndicator() {
	for i := 0; i < N_FLOORS; i++ {
		if ElevGetFloorSensorSignal() == i{
			ElevSetFloorIndicator(i)
		}
		
	}
	
}