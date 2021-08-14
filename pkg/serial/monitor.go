package serial

import (
	"fmt"
	"github.com/tarm/serial"
	"strconv"
	"strings"
	"time"
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

type ZQI22Message struct {
	SourceAspectRatio uint8
	SourceFrameRate uint8
	SourceVerticalResolution uint16
	HDR bool
}

// StartZQI22Monitor begins reading the serial connection. When a complete
// ZQI22 message is found, it calls the provided function.
// TODO: add cancel channel to break out of monitor loop
func (q *LumagenSession) StartZQI22Monitor(callback func(command ZQI22Message)) error {
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
					cmd, err := parseZQI22(msg)
					if err != nil {
						fmt.Printf("Error parsing line '%s' due to '%v', ignoring\n", msg, err)
						continue
					}
					callback(cmd)
				} else {
					line = append(line, buf[i])
				}
			}
		}
	}()

	return nil
}

func parseZQI22(line string) (msg ZQI22Message, err error) {
	line = strings.TrimSpace(line)
	if !strings.HasPrefix(line, MessagePrefixI22) {
		err = fmt.Errorf("message '%s' is not recognized as type %s", line, MessagePrefixI22)
		return
	}

	parts := strings.Split(line, ",")
	value, err := strconv.ParseUint(parts[I22IndexSourceFrameRate], 10, 8)
	if err != nil {
		return
	}

	msg.SourceFrameRate = uint8(value)

	value, err = strconv.ParseUint(parts[I22IndexSourceVResolution], 10, 16)
	if err != nil {
		return
	}

	msg.SourceVerticalResolution = uint16(value)

	value, err = strconv.ParseUint(parts[I22IndexSourceAspectRatio], 10, 8)
	if err != nil {
		return
	}

	msg.SourceAspectRatio = uint8(value)

	msg.HDR, err = strconv.ParseBool(parts[I22IndexHDR])
	if err != nil {
		return
	}

	return
}