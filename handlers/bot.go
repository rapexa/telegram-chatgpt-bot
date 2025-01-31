package handlers

import (
	"log"
	"telegram-chatgpt-bot.com/m/config"
	"telegram-chatgpt-bot.com/m/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartBot initializes the Telegram bot
func StartBot(logger *log.Logger) error {
	bot, err := tgbotapi.NewBotAPI(config.GetEnv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		return err
	}

	bot.Debug = false
	logger.Println("Telegram bot started")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go handleMessage(update.Message, bot, logger)
	}
	return nil
}

func handleMessage(msg *tgbotapi.Message, bot *tgbotapi.BotAPI, logger *log.Logger) {
	response, err := services.GetChatGPTResponse(msg.Text)
	if err != nil {
		logger.Println("Error getting ChatGPT response:", err)
		response = "Sorry, I couldn't process your request."
	}

	reply := response
	reply += "\n\n[Sponsored]: rapexa"

	// Send response
	msgConfig := tgbotapi.NewMessage(msg.Chat.ID, reply)
	bot.Send(msgConfig)
}
