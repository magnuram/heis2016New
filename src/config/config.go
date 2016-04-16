package config

import (
	"os/exec"
)

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
	ON  bool = true
	OFF bool = false

	//LAMP CALL
	BtnUp     int = 0
	BtnDown   int = 1
	BtnInside int = 2

	//STATES

	IDLE     int = 0
	MOVING   int = 1
	DOOROPEN int = 2
)

// Start a new terminal when restart.Run()
var Restart = exec.Command("gnome-terminal", "-x", "sh", "-c", "go run main.go")

var CloseConnectionChan = make(chan bool)

type ButtonPress struct {
	Button int
	Floor  int
}

type ELINFO struct {
	State    int
	Floor    int
	Dir      int
	ReqUp    [N_FLOORS]int
	ReqDown  [N_FLOORS]int
	ReqLocal [N_FLOORS]int
}

var Elinfo = ELINFO{}

// Network message
type Message struct {
	Category int
	Floor    int
	Button   int
	Cost     int
	Addr     string `json:"-"`
}
