package main

import (
	"machine"
	"time"
)

var out = machine.PinConfig{Mode: machine.PinOutput}

func main() {
	led := machine.D1
	led.Configure(out)
	for {
		led.High()
		time.Sleep(2 * time.Second)
		led.Low()
		time.Sleep(2 * time.Second)
	}
}
