package main

import (
	"github.com/sirupsen/logrus"
	"telegram-chatgpt-bot.com/m/config"
	handler "telegram-chatgpt-bot.com/m/handlers"
)

func main() {
	// Initialize Logrus
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)

	// Load configuration
	config.LoadEnv()

	// Start the bot
	handler.StartBot()
}
