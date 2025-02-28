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
*
 */
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/SDBlackwood/hexedit/internal"
)

func main() {
	var logger *slog.Logger
	logger = internal.Logger()

	// Parse the command line arguments
	filePath := flag.String("filepath", "", "Path to the the input file")

	// Parse the command line variables
	flag.Parse()

	logger.Debug("cmd args", "filepath", *filePath)

	if *filePath == "" {
		fmt.Println("File path cannot be empty")
		logger.Error("Empty cmd args", "filepath", *filePath)
		os.Exit(1)
	}

	// Read the file into memory
	fileBytes, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Cannot read file path %v\n", *filePath)
		logger.Error("invalid filepath", "filepath", *filePath)
	}

	// Convert byte data to hex
	// A bytes is a 8bit value
	for i := 0; i < len(fileBytes); i++ {
		fmt.Printf("%v:%v:%v", string(fileBytes[i]), fileBytes[i], internal.HexConvert(fileBytes[i]))		
	}
}
