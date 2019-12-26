package touchscreen

import (
	"machine"
	"time"
)

type Pointer interface {
	GetTouchPoint() Point
}

type Point struct {
	X int
	Y int
	Z int
}

var (
	input  = machine.PinConfig{Mode: machine.PinInput}
	output = machine.PinConfig{Mode: machine.PinOutput}
)

func delayMicroseconds(n int) {
	interval := int64(time.Duration(n) * time.Microsecond)
	for now := time.Now().UnixNano(); time.Now().UnixNano() < (now + interval); {
	}
}
