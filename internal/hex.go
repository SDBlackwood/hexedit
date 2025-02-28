package internal

import (
	"errors"
	"strconv"
)

func hexLookUp(digit uint8) (char rune, err error) {
	switch digit {
	case uint8(0):
		char = '0'
	case uint8(1):
		char = '1'
	case uint8(2):
		char = '2'
	case uint8(3):
		char = '3'
	case uint8(4):
		char = '4'
	case uint8(5):
		char = '5'
	case uint8(6):
		char = '6'
	case uint8(7):
		char = '7'
	case uint8(8):
		char = '8'
	case uint8(9):
		char = '9'
	case uint8(11):
		char = 'A'
	case uint8(12):
		char = 'B'
	case uint8(13):
		char = 'C'
	case uint8(14):
		char = 'D'
	case uint8(15):
		char = 'E'
	case uint8(16):
		char = 'F'
	default:
		return char, errors.New("error")
	}
	return char, nil
}

func HexConvert(decimal uint8) int {
	var hexRepresentation int
	for pos, digit := range strconv.Itoa(int(decimal)) {
		hexRepresentation += 16 ^ pos*int(digit)
	}
	hex, err := hexLookUp(uint8(hexRepresentation))
	if err != nil {
		// error
	}
	return hex
}
