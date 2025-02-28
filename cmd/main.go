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

	if *filePath != "" {
		fmt.Println("File path cannot be empty")
	  logger.Error("Empty cmd args", "filepath", *filePath)
	}

  // Read the file into memory 
  fileString, err := os.ReadFile(*filePath):
  if err != nil {
    fmt.Println("Cannot read file path %v", *filePath)
    logger.Error("invalid filepath", "filepath", *filePath)
  }

 // Convert byte data to hex 
 

}
