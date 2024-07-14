package config

import (
	"log"
	"os"
)

type Config struct {
	OpenAIKey string
}

func LoadConfig() *Config {
	openAIKey := os.Getenv("OPENAI_API_KEY")
	if openAIKey == "" {
		log.Fatal("OPENAI_API_KEY not set in environment variables")
	}

	return &Config{
		OpenAIKey: openAIKey,
	}
}
