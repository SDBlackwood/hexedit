package representation

import (
	"fmt"
	"log/slog"
	"unicode/utf8"
)

func Render(logger *slog.Logger, line string) []string {
	// Render applies all of the representations onto the terminal, taking into account the width.
	// Each character vertically aligns with the first hexideximal code point
	// We need to determine how many hex values are associated with each character
	// To do this we need to iterate over each character and determine its utf-8 value
	// and the add them to 2 congruent arrays: 1 for raw text, 1 for hex

	// Append everything to a slice which can then be rendered onto the terminal
	// This allows us to verify what is in the output.
	// NOTE: we can chunk this up at some point so that we don't have to hold it in memory
	var textBuf []rune
	var hexBuf []rune
	for _, char := range line {
		// Determine the number of bytes in the character
		byteLength := utf8.RuneLen(char)

		logger.Debug("info", "byteLength", byteLength, "char", string(char))

		// Append the character to the padded line buffer with the correct amount of padding
		paddedChar := fmt.Sprintf("%-*s", byteLength*3, string(char))

		// Append the padded character to the text buffer
		textBuf = append(textBuf, []rune(paddedChar)...)

		// Append the hex value of the character to the hex buffer
		hexBuf = append(hexBuf, []rune(Hex(uint8(char)))...)
		hexBuf = append(hexBuf, ' ')
	}

	// Return the text and hex buffers as a slice of strings
	// This allows us to print the text and hex side by side
	return []string{string(textBuf), string(hexBuf)}
}
