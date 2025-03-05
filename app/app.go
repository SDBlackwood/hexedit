package app

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/SDBlackwood/hexedit/representation"
	"golang.org/x/term"
)

type App struct {
	filePath     string
	fileHandler  *os.File
	logger       *slog.Logger
	tW, tH       int
	closeChannel chan int
}

// TUIApp is the interface for the App struct
type TUIApp interface {
	OpenFile() (err error)
	Close() (err error)
	Run()
	HandleEvents() (err error)
	Render() (err error)
}

func NewApp(filePath string, logger *slog.Logger) *App {
	if filePath == "" {
		fmt.Println("File path cannot be empty")
		logger.Error("Empty cmd args", "filepath", filePath)
		os.Exit(1)
	}

	// Get the terminal size
	tW, tH, err := term.GetSize(0)
	if err != nil {
		fmt.Println("Error getting terminal size")
		logger.Error("error getting terminal size", "error", err)
	}

	closeChannel := make(chan int, 1)

	return &App{
		filePath:     filePath,
		logger:       logger,
		tW:           tW,
		tH:           tH,
		closeChannel: closeChannel,
	}
}

func (a *App) OpenFile() (err error) {
	a.fileHandler, err = os.Open(a.filePath)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Close() (err error) {
	err = a.fileHandler.Close()
	if err != nil {
		return err
	}
	close(a.closeChannel)
	return nil
}

func (a *App) Run() {
	// Run starts an event loop which handles events and renders the UI
	// Make a channel for keyboard interuptes
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for {
		go func() {
			for {
				select {
				case <-c:
				case <-a.closeChannel:
				}
				a.Close()
				os.Exit(0)
			}
		}()
	}
}

func (a *App) HandleEvents() (err error) {
	//}/ HandleEvents handles events from the event loop
	return nil
}

func (a *App) Render(pipeOutput bool) (err error) {
	// Render renders the UI on the initialisation of the app
	// Read line by line in the file handler and send to a channel
	// If we are outputing and exiting, send a signal to the event loop to close
	scanner := bufio.NewScanner(a.fileHandler)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			rendered := representation.Render(a.logger, line)
			for _, line := range rendered {
				fmt.Println(line)
			}
		}
		if pipeOutput {
			a.closeChannel <- 1
		}
	}()

	return nil
}
