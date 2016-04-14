package config

const (
	//QUANTITY
	N_FLOORS    int = 4
	N_ELEVATORS int = 1
	N_BUTTONS   int = 3

	//DIRECTIONS
	UP_DIR   int = 1
	DOWN_DIR int = -1
	STOP_DIR int = 0

	//LIGHTS
	ON  int = 1
	OFF int = 0

	//LAMP CALL
	BUTTON_CALL_UP   int = 0
	BUTTON_CALL_DOWN int = 1
	BUTTON_COMMAND   int = 2

	//STATES
	INIT      int = 0
	IDLE      int = 1
	MOVING    int = 2
	DOOROPEN int = 3
	STOP      int = 4

	//ELEVATOR TYPES
	ELEVTYPE_COMEDI     int = 0
	ELEVTYPE_SIMULATION int = 1
)

type ELINFO struct {
	State     int
	PrevFloor int
	Dir       int
	ReqUp     [N_FLOORS]int
	ReqDown   [N_FLOORS]int
	ReqLocal  [N_FLOORS]int
}

var Elinfo = ELINFO{}
