package fsm

import (
	. "../config"
	//. "../queue"
	//. "../elev"
	"log"
	//"time"
)

var state int
var floor int
var dir int

type Channels struct {
	//Events
	NewOrder     chan bool
	FloorReached chan int
	doorTimeout  chan bool
	//Hardware interaction
	MotorDir  chan int
	FloorLamp chan int
	DoorLamp  chan bool
	//Door timer
	doorTimerReset chan bool
	//Network Interaction
	OutgoingMsg chan Message
}

func FsmInit(ch Channels, startFloor int) {
	state = IDLE
	dir = STOP_DIR
	floor = startFloor

}

func run(ch Channels) {
	for {
		select {
		case <-ch.NewOrder:
			eventNewOrder(ch)

		case floor := <-ch.FloorReached:
			eventFloorReached(ch, floor)
			//case <-ch.doorTimeout:
		}
	}
}

func eventNewOrder(ch Channels) {
	switch state {
	case IDLE:
		ch.DoorLamp <- true
		state = DOOROPEN
	case MOVING:

	case DOOROPEN:
		CloseConnectionChan <- true

	default:
		CloseConnectionChan <- true
		Restart.Run()
	}
}

func eventFloorReached(ch Channels, newFloor int) {
	log.Printf("%sEVENT: Floor %d reached in state", newFloor+1)
	floor = newFloor
	ch.FloorLamp <- floor
	switch state {
	case MOVING:
		ch.DoorLamp <- true
		dir = STOP_DIR
		ch.MotorDir <- dir
		state = DOOROPEN

	default:
		CloseConnectionChan <- true
		Restart.Run()

	}
}
