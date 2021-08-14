package main

import (
	"fmt"
	"github.com/mlsorensen/lumagen/pkg/serial"
)

func main() {
	mon := serial.LumagenSession{SerialPort: "/dev/ttyUSB1"}
	err := mon.StartZQI22Monitor(handleMessage)
	if err != nil {
		panic(err)
	}

	select{}
}

func handleMessage(command serial.ZQI22Message) {
	fmt.Printf("got this: %v\n", command)
}