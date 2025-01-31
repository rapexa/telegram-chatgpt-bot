package services

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"telegram-chatgpt-bot.com/m/config"
)

// GetChatGPTResponse sends a user's question to ChatGPT API and returns the response
func GetChatGPTResponse(question string) string {
	client := openai.NewClient(config.GetEnv("OPENAI_API_KEY"))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo, // Use GPT-3.5 Turbo model
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: "You are an AI assistant answering coding questions."},
				{Role: "user", Content: question},
			},
		},
	)

	if err != nil {
		logrus.WithError(err).Error("Failed to get response from ChatGPT API")
		return "Sorry, I couldn't process your request."
	}

	return resp.Choices[0].Message.Content
}
