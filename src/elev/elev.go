package elev

import (
	//"C"
	. "../channels"
	. "../config"
	. "../driver"
	"errors"
	"log"
	"time"
)

/*
const (
	N_FLOORS  = 4 //Number of floors, Hardware-depentent
	N_BUTTONS = 3 //Number of buttons/lamps on a per-floor basis
)
*/
var lampChannelMatrix = [N_FLOORS][N_BUTTONS]int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var buttonChannelMatrix = [N_FLOORS][N_BUTTONS]int{
	{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
	{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
	{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
	{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

const (
	MOTOR_SPEED = 2800
)
const elevStopDelay = 150 * time.Millisecond

/*
const (
	UP   = 1
	STOP = 0
	DOWN = -1
)
*/
//Initialises lift and moves it to a defined state
func ElevInit() (int, error) {
	//Init Io hardware
	if !IoInit() {
		return -1, errors.New("Elev: IoInit() failed")
	}
	//Clear all floor lamps
	for i := 0; i < N_FLOORS; i++ {
		if i != 0 {
			SetButtonLamp(BtnDown, i, OFF)
		}
		if i != N_FLOORS-1 {
			SetButtonLamp(BtnUp, i, OFF)
		}
		SetButtonLamp(BtnInside, i, OFF)
	}
	//Clear stop and door open lamp
	SetStopLamp(OFF)
	SetDoorLamp(OFF)

	//Move to defined state
	for ReadFloor() == -1 {
		SetMotorDir(DOWN_DIR)
	}
	time.Sleep(elevStopDelay)
	SetMotorDir(STOP_DIR)
	SetFloorLamp(ReadFloor())
	var floor = ReadFloor()
	return floor, nil
}

func SetMotorDir(direction int) {
	if direction == 0 {
		IoWriteAnalog(MOTOR, 0)
	} else if direction > 0 {
		IoClearBit(MOTORDIR)
		IoWriteAnalog(MOTOR, MOTOR_SPEED)
	} else if direction < 0 {
		IoSetBit(MOTORDIR)
		IoWriteAnalog(MOTOR, MOTOR_SPEED)
	}
}

func SetFloorLamp(floor int) {
	if floor < 0 || floor >= N_FLOORS {
		log.Println("Error: Floor %d out of range\n", floor)
		log.Println("No floor indicator set")
		return
	}
	//Binary encoding. One light always be on
	if floor&0x02 > 0 {
		IoSetBit(LIGHT_FLOOR_IND1)
	} else {
		IoClearBit(LIGHT_FLOOR_IND1)
	}

	if floor&0x01 > 0 {
		IoSetBit(LIGHT_FLOOR_IND2)
	} else {
		IoClearBit(LIGHT_FLOOR_IND2)
	}
}
func ReadFloor() int {
	if IoReadBit(SENSOR_FLOOR1) {
		return 0
	} else if IoReadBit(SENSOR_FLOOR2) {
		return 1
	} else if IoReadBit(SENSOR_FLOOR3) {
		return 2
	} else if IoReadBit(SENSOR_FLOOR4) {
		return 3
	} else {
		return -1
	}
}

func SetDoorLamp(value bool) {
	if value {
		IoSetBit(LIGHT_DOOR_OPEN)
	} else {
		IoClearBit(LIGHT_DOOR_OPEN)
	}
}

func SetButtonLamp(button int, floor int, value bool) {
	if floor < 0 || floor >= N_FLOORS {
		log.Println("Error: Floor %d out of range\n", floor)
	}
	if button == BtnUp && floor == N_FLOORS-1 {
		log.Println("Button up from top floor does not exist")
		return
	}
	if button == BtnDown && floor == 0 {
		log.Println("Button down from first floor does not exist!")
		return
	}
	if button != BtnUp && button != BtnDown && button != BtnInside {
		log.Printf("Invalid button %d\n", button)
		return
	}
	if value {
		IoSetBit(lampChannelMatrix[floor][button])
	} else {
		IoClearBit(lampChannelMatrix[floor][button])
	}
}

func ReadButton(button int, floor int) bool {
	//Different error messages
	if floor < 0 || floor >= N_FLOORS {
		log.Println("Error: Floor %d out of range\n", floor)
		return false
	}
	if button < 0 || button >= N_BUTTONS {
		log.Println("Error: Button %d out of range\n", button)
		return false
	}
	if button == BtnDown && floor == 0 {
		log.Println("Button down from first floor does not exist")
		return false
	}
	if button == BtnUp && floor == N_FLOORS-1 {
		log.Println("Button up from top floor does not exist")
		return false
	}

	if IoReadBit(buttonChannelMatrix[floor][button]) == true {
		return true
	} else {
		return false
	}
}

//----------------DOESNT USE YET --------------//
func GetStopSignal() bool {
	if IoReadBit(STOP_BUTTON) {
		return true
	} else {
		return false
	}
}

func SetStopLamp(value bool) {
	if value {
		IoSetBit(LIGHT_STOP)
	} else {
		IoClearBit(LIGHT_STOP)
	}
}

func SGetObstructionSignal() bool {
	return IoReadBit(OBSTRUCTION)
}
