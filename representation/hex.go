package representation

import (
	"strconv"
	"strings"
)

// Hex converts a decimal value to a 2 char hexidecimal value as a string
func Hex(decimal uint8) (hexRepresentation string) {
	// Convert the decimal value to a 2 char hexidecimal value as a string
	hexRepresentation = strconv.FormatInt(int64(decimal), 16)
	if len(hexRepresentation) == 1 {
		hexRepresentation = "0" + hexRepresentation
	}
	hexRepresentation = strings.ToUpper(hexRepresentation)
	return hexRepresentation
}
