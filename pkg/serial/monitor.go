package serial

import (
	"fmt"
	"github.com/mlsorensen/lumagen/pkg/serial/parsers"
	"github.com/tarm/serial"
	"strings"
	"time"
)

const(
	DefaultBaud = 9600
)

type LumagenSession struct {
	SerialPort string
	openPort *serial.Port
}

func (l *LumagenSession) NewSession() error {
	c := &serial.Config{Name: l.SerialPort, Baud: DefaultBaud}
	s, err := serial.OpenPort(c)
	if err != nil {
		return err
	}
	l.openPort = s
	return nil
}

// StartMessageMonitor begins reading the serial connection. When a complete
// message is found, it calls the provided parser functions which can attempt to
// parse and handle the message as they see fit.
// TODO: add cancel channel to break out of monitor loop?
func (q *LumagenSession) StartMessageMonitor(parsers []parsers.Parser) error {
	if q.openPort == nil {
		err := q.NewSession()
		if err != nil {
			return err
		}
	}

	var line []byte
	buf := make([]byte, 128)

	go func() {
		for {
			num, err := q.openPort.Read(buf)
			if err != nil {
				// log error, throttle retries to 10s
				time.Sleep(time.Second * 10)
				continue
			}

			for i := 0; i < num; i++ {
				if buf[i] == '\r' {
					// log string and error
					msg := strings.TrimSpace(string(line))
					line = []byte{}
					for _, parser := range parsers {
						err := parser.Parse(msg)
						if err != nil {
							fmt.Printf("Error parsing line '%s' due to\"%v\", ignoring\n", msg, err)
							continue
						}
					}
				} else {
					line = append(line, buf[i])
				}
			}
		}
	}()

	return nil
}