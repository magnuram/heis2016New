package main

import (
	. "./src/config"
	. "./src/elev"
	. "./src/fsm"
	"log"
	"time"
)

func main() {
	var floor int
	var err error
	floor, err = ElevInit()
	if err != nil {
		Restart.Run() //Restarts main if failed to initialise
		log.Println("ERROR -> Main: \t Hardware init failure")
		log.Fatal(err)
	} else {
		log.Println("Hardware init complete")
	}

	//Make channels
	ch := Channels{
		NewOrder:     make(chan bool),
		FloorReached: make(chan int),
		//Hardware interaction
		MotorDir:  make(chan int, 10), //buffersize 10
		FloorLamp: make(chan int, 10),
		DoorLamp:  make(chan bool, 10),
		//Network Interaction
		//OutgoingMsg make(chan Message)
	}

	FsmInit(ch, floor)
	log.Println("Fsm Init complete")
}

//go eventHandler(ch)
//buttonchan := pollButtons()
/*
	for {
		select {
		case key := <-buttonchan:
			switch key.Button {
			case BtnInside:


			case BtnUp, BtnDown:
				log.Println("ok2")
			}
		}
	}

}
*/
/*
func eventHandler() {
	buttonChan := pollButtons()
	floorChan := pollFloors()

	for {
		select {
		case key := <-buttonChan:
			switch key.Button {
			case BtnInside:

			case BtnUp, BtnDown:

			}

		case floor := <-floorChan:
			ch.FloorReached <- floor
		case dir := <-ch.MotorDir:
			SetMotorDir(dir)
		case floorlamp := <-ch.FloorLamp:
			SetFloorLamp(floorlamp)
		case value := <-ch.DoorLamp:
			SetDoorLamp(value)
		}
	}
}
*/

func pollButtons() <-chan ButtonPress {
	c := make(chan ButtonPress)
	go func() {
		var buttonState [N_FLOORS][N_BUTTONS]bool

		for {
			for f := 0; f < N_FLOORS; f++ {
				for b := 0; b < N_BUTTONS; b++ {
					if (f == 0 && b == BtnDown) ||
						(f == N_FLOORS-1 && b == BtnUp) {
						continue
					}
					if ReadButton(f, b) {
						if !buttonState[f][b] {
							c <- ButtonPress{Button: b, Floor: f}
						}
						buttonState[f][b] = true
					} else {
						buttonState[f][b] = false
					}
				}
			}
			time.Sleep(time.Millisecond)
		}
	}()
	return c
}

func pollFloors() <-chan int {
	c := make(chan int)
	go func() {
		oldFloor := ReadFloor()

		for {
			newFloor := ReadFloor()
			if newFloor != oldFloor && newFloor != -1 {
				c <- newFloor
			}
			oldFloor = newFloor
			time.Sleep(time.Millisecond)
		}
	}()
	return c
}
