package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"telegram-chatgpt-bot.com/m/appStrings"
	"telegram-chatgpt-bot.com/m/config"
	"telegram-chatgpt-bot.com/m/services"
)

// startBot initializes and starts the Telegram bot
func startBot() {
	bot, err := tgbotapi.NewBotAPI(config.GetEnv("BOT_TOKEN"))
	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize Telegram bot")
	}

	bot.Debug = true
	logrus.Infof("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		logrus.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message == nil {
			continue
		}

		var response string

		if update.Message.Text == "/start" {
			response = appStrings.StartString
		} else {
			response = services.GetChatGPTResponse(update.Message.Text) // Get response from ChatGPT
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		_, err := bot.Send(msg)
		if err != nil {
			logrus.WithError(err).Error("Failed to send message")
		}
	}
}
