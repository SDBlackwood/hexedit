package internal

import (
	"fmt"
	"log/slog"
	"os"
)

var logger *slog.Logger

func init() {
	f, err := os.Create("./debug.log")
	if err != nil {
		fmt.Println("Error Creating log")
	}
	// Create the logger and set the loglevel
	logger = slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}

func Logger() *slog.Logger {
	return logger
}
