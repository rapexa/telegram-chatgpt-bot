package main

import (
	"telegram-chatgpt-bot.com/m/config"
	"telegram-chatgpt-bot.com/m/handlers"
	"telegram-chatgpt-bot.com/m/utils"
)

func main() {
	// Load configuration
	config.LoadEnv()

	// Initialize logger
	logger := utils.NewLogger()

	// Start Telegram bot
	err := handlers.StartBot(logger)
	if err != nil {
		logger.Fatal("Failed to start bot:", err)
	}
}
