package driver

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/
import "C"

func IoInit() bool {
	return int(C.io_init()) != 0
}

func IoSetBit(channel int) {
	C.io_set_bit(C.int(channel))
}

func IoClearBit(channel int) {
	C.io_clear_bit(C.int(channel))
}

func IoWriteAnalog(channel int, value int) {
	C.io_write_analog(C.int(channel), C.int(value))
}

func IoReadBit(channel int) bool {
	return bool(int(C.io_read_bit(C.int(channel))) != 0)
}

func IoReadAnalog(channel int) int {
	return int(C.io_read_analog(C.int(channel)))
}
