package queue

import (
	//. "../driver"
	. "../config"
	. "../elev"
)

func QOrderExist() bool {
	for i := 0; i < N_FLOORS; i++ {
		if Elinf.ReqLocal[i] != 0 || Elinf.ReqUp[i] != 0 || Elinf.ReqDown[i] != 0 {
			return true
		}
	}
	return false
}

func QSetLight(){
	for i :=0 ; i < N_FLOORS; i++{
		ElevSetButtonLamp(BUTTON_COMMAND,i,Elinfo.ReqLocal[i])
		if i != 0{
		ElevSetButtonLamp(BUTTON_CALL_DOWN,i,Elinfo.ReqDown[i])
		}
		if i != 3{
		ElevSetButtonLamp(BUTTON_CALL_UP,i,Elinfo.ReqUp[i])
		}
	}

}

func QAddOrder(floor int, btnTypePres int) bool{

	select{
		switch btnTypePres{
		case 0: //button call up 
			Elinf.ReqUp[floor] = 1
			break
		case 1:
			Elinf.ReqDown[floor] = 1
			break
		case 2:
			Elinf.ReqLocal[floor] = 1
			break

	}
}
}

func QOrdersAbove(currentFloor int) bool{
		for flr := currentFloor+1;  flr<N_FLOORS; flr++ {
		if Elinf.ReqLocal[flr] != 0 || Elinf.ReqUp[flr] != 0 || Elinf.ReqDown[flr] != 0{
			return true
		} 
	}
	return false
}

func QOrdersBelow(currentFloor int) bool {
	for flr := 0;  flr<currentFloor; flr++ {
		if Elinf.ReqLocal[flr] != 0 || Elinf.ReqUp[flr] != 0 || Elinf.ReqDown[flr] != 0{
			return true
		} 
	}
	return false
	
}

func QChooseDir(currentFloor int, prevDir int)int{
	select{
		switch prevDir{
		case UP_DIR:
			if QOrdersAbove(currentFloor){
					return UP_DIR
				}else if QOrdersBelow(currentFloor){
					return DOWN_DIR
				}else{
					return STOP_DIR
				}

		case DOWN_DIR:
			if QOrdersAbove(currentFloor){
					return DOWN_DIR
				}else if QOrdersBelow(currentFloor){
					return UP_DIR
				}else{
					return STOP_DIR
				}

		case STOP_DIR:
			if QOrdersAbove(currentFloor){
					return UP_DIR
				}else if QOrdersBelow(currentFloor){
					return DOWN_DIR
				}else{
					return STOP_DIR
				}
		default:
			return STOP_DIR
		}
	}
}

func QShouldStop(flr int, prevDir int)int {
	if prevDir == -1{
		if Elinf.ReqDown[flr] != 0  || Elinf.ReqLocal[flr] != 0 || !QOrdersBelow(flr) || flr == 0{
			return 0
		}
	}
	if prevDir == 1{
		if Elinf.ReqUp[flr] != 0 || Elinf.ReqLocal[flr] != 0 || !QOrdersAbove(flr) || flr == 3{
			return 0
		}
	}
	return 1
}

func QDeleteOrders() {
	for i := 0; i < N_FLOORS; i++ {
		Elinf.ReqUP[i] 		= 0
		Elinf.ReqDown[i] 	= 0
		Elinf.ReqLocal[i]	= 0
	}
	QSetLight()
	
}

func QDeleteManual(flr int){
		Elinf.ReqUP[flr] 	= 0
		Elinf.ReqDown[flr] 	= 0
		Elinf.ReqLocal[flr] = 0
		QSetLight()
}