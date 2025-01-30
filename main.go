package main

import (
	"telegram-chatgpt-bot.com/m/config"
	"telegram-chatgpt-bot.com/m/database"
	"telegram-chatgpt-bot.com/m/utils"
)

func main() {

	// Load configuration
	config.LoadEnv()

	// Initialize logger
	logger := utils.NewLogger()

	// Connect to database
	db, err := database.ConnectDB()
	if err != nil {
		logger.Fatal("Database connection failed:", err)
	}
	defer db.Close()

}
