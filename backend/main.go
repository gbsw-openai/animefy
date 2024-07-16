package main

import (
	"backend/controller"
	"backend/repository"
	"backend/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey == "" {
		log.Fatal("OpenAI API key not found in environment variables")
	}

	repo := repository.NewInMemoryAnimeRepository()
	jsonFilePath := "./anime.json"
	if err := repo.LoadAnimesFromJSON(jsonFilePath); err != nil {
		log.Fatalf("Failed to load anime data from JSON file: %v", err)
	}

	chatController := controller.NewChatController(repo, openaiKey)

	r := routes.SetupRouter(chatController)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
