package main

import (
	."./src/elev"
	."./src/queue"
	."./src/config"'
	."./src/fsm'
)

func main() {
	if err := ElevInit(); err != nil {
		log.Println("ERROR -> Main: \t Hardware init failure")
		log.Fatal(err)
	} else {
		log.Println("Hardware init complete")
	}
	FsmInit()

	for{

		var prevButton[N_FLOORS][N_BUTTONS]int
		for flr := 0; flr < N_FLOORS; flr++ {
			for btn := 0; btn < N_BUTTONS; btn++ {
				var button int = ElevGetButtonSignal(btn,flr)
				if prevButton[flr][btn] != button && button == 1{
					FsmButtonPressed(btn,flr)
				}
				prevButton[flr][btn] = button
			}
			
		}

		QSetLight()
		FsmSetIndicator()

		var prevOrderExist bool
		var Order bool = QOrderExist()
		if Order != prevOrderExist{
			if Order == true{
				FsmOrderExist()
			}
		}
		prevOrderExist = Order

		var prevCorrectFloorReached int = -1
		var floor int = ElevGetFloorSensorSignal()
		if floor != prevCorrectFloorReached && floor != -1{
			FsmCorrectFloorReached(floor)
		}
		prevCorrectFloorReached(floor)

		var prevTimeOut int
		var timer int = 3
		if timer != prevTimeOut{
			if timer == 3 {
				FsmTimeOut()
			}
		}
		prevTimeOut = timer

	}

}
