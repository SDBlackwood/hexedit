/*
* TUI Application which takes a file path and represents the data in a hexviewer (plus some other char sets)
*
* When we parse a file we will get a slice of bytes i.e. []byte.
* Each byte is an alias for a uint8, an unsigned 8 bit integer.
* So for example 's' is 115 and 'T' is 84.
*
* The uint8 has values between 0-255 and maps directly on
*
* Note the first 8 bits / 128 values of ASCII map directly onto UTF-8
*
* ## Hexidecimal
* We can calculate the hexidecimal by iterating each value of the uint8 values and converting as follows:
* For each number in the decimal value e.g. 115 -> [1,1,5] is the sume of each digit i.e 16^i where i is the index.
* or we could just have a look up...
* ## Octal
* Octal is a numeral system that uses the base 8, and the numbers 0 through 7.
* It's also known as base 8 or octonary.
*
* We want to represent the text in the following ways:
* - Hexidecimal
* - Text (ASCII)
* - binary
* - octal
* - uint8
* - int8
* - uint16
* - int16
* - uint32
* - int32
* - uint64
* - int64
* - ULEB128
* - SLEB128
* - float16
* - bfloat16
* - float32
* - float64
* - GUID
* - ASCII
* - UTF-8
* - UTF-16
* - GB18030
* - BIG5
* - SHIFT-JIS
 */
package main

import (
	"flag"
	"os"

	"github.com/SDBlackwood/hexedit/app"
	"github.com/SDBlackwood/hexedit/internal"
)

func main() {
	logger := internal.Logger()

	// Parse the command line arguments
	filePath := flag.String("f", "", "Path to the the input file")
	pipeOutput := flag.Bool("o", false, "Output the contents and exit")

	// Parse the command line variables
	flag.Parse()

	app := app.NewApp(*filePath, logger)
	err := app.OpenFile()

	defer app.Close()

	if err != nil {
		logger.Error("error opening file", "error", err)
		os.Exit(1)
	}

	logger.Debug("cmd args", "filePath", *filePath)

	// Render the UI
	err = app.Render(*pipeOutput)
	if err != nil {
		logger.Error("error rendering UI", "error", err)
		os.Exit(1)
	}

	// Start the event loop
	app.Run()

	// Handle events/
}
