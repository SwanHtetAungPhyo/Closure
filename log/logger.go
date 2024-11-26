package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	green = "\033[32m"
	reset = "\033[0m"
)



var Logger *log.Logger

func init() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	Logger = log.New(multiWriter, fmt.Sprintf("%s[CLOSURE] %s", "\033[32m", "\033[0m"), log.LstdFlags)
}
