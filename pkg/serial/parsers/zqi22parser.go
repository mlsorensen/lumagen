package parsers

import (
	"fmt"
	"github.com/mlsorensen/lumagen/pkg/serial/message"
	"strconv"
	"strings"
)

const (
	MessageStartToken = "!"
	MessagePrefixI22 = "!I22"
	I22FieldCount = 20

	I22IndexSourceFrameRate = 2
	I22IndexSourceVResolution = 3
	I22IndexSourceAspectRatio = 7
	I22IndexHDR = 17
)

type ZQI22Parser struct {
	Handler func(message message.ZQI22Message)
}

func (z22 ZQI22Parser) Parse(line string) error {
	var msg message.ZQI22Message
	line = extractRightByDelimiter(strings.TrimSpace(line), MessageStartToken)
	if !strings.HasPrefix(line, MessagePrefixI22) {
		return fmt.Errorf("message '%s' is not recognized as type %s", line, MessagePrefixI22)
	}

	parts := strings.Split(line, ",")
	value, err := strconv.ParseUint(parts[I22IndexSourceFrameRate], 10, 8)
	if err != nil {
		return err
	}

	msg.SourceFrameRate = uint8(value)

	value, err = strconv.ParseUint(parts[I22IndexSourceVResolution], 10, 16)
	if err != nil {
		return err
	}

	msg.SourceVerticalResolution = uint16(value)

	value, err = strconv.ParseUint(parts[I22IndexSourceAspectRatio], 10, 8)
	if err != nil {
		return err
	}

	msg.SourceAspectRatio = uint8(value)

	msg.HDR, err = strconv.ParseBool(parts[I22IndexHDR])
	if err != nil {
		return err
	}

	if z22.Handler != nil {
		z22.Handler(msg)
	}

	return nil
}

// extractRightByDelimiter returns empty string if delimiter is not found, otherwise returns substring after
// delimiter
func extractRightByDelimiter(msg, delimiter string) string {
	idx := strings.Index(msg, delimiter)
	if idx < 0 {
		return ""
	}

	return msg[idx:]
}
