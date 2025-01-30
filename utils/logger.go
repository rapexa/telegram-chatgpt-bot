package utils

import (
	"log"
	"os"
)

// NewLogger initializes a new logger
func NewLogger() *log.Logger {
	file, err := os.OpenFile("bot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	return log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
