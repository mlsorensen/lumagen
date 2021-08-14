package main

import (
	"fmt"
	"github.com/mlsorensen/lumagen/pkg/serial"
	"github.com/mlsorensen/lumagen/pkg/serial/message"
	"github.com/mlsorensen/lumagen/pkg/serial/parsers"
)

func main() {
	mon := serial.LumagenSession{SerialPort: "/dev/ttyUSB1"}
	parser := parsers.ZQI22Parser{Handler: handleZQI22Message}
	err := mon.StartMessageMonitor([]parsers.Parser{parser})
	if err != nil {
		panic(err)
	}

	select{}
}

func handleZQI22Message(msg message.ZQI22Message) {
	fmt.Printf("got this: %v\n", msg)
}