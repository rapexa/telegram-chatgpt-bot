package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment")
	}
}

// GetEnv retrieves environment variables
func GetEnv(key string) string {
	return os.Getenv(key)
}
