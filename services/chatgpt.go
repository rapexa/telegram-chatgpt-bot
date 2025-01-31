package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"telegram-chatgpt-bot.com/m/config"
)

// ChatGPTRequest represents OpenAI request structure
type ChatGPTRequest struct {
	Model     string `json:"model"`
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

// ChatGPTResponse represents OpenAI response structure
type ChatGPTResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// GetChatGPTResponse sends a request to OpenAI API
func GetChatGPTResponse(prompt string) (string, error) {
	apiKey := config.GetEnv("OPENAI_API_KEY")
	url := "https://api.openai.com/v1/completions"

	requestData := ChatGPTRequest{
		Model:     "text-davinci-003",
		Prompt:    prompt,
		MaxTokens: 150,
	}

	body, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response ChatGPTResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if len(response.Choices) > 0 {
		return response.Choices[0].Text, nil
	}
	return "No response from AI", nil
}
