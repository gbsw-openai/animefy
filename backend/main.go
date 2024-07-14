package main

import (
	controller "backend/controller"
	"backend/repository"
	service "backend/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repo := repository.NewInMemoryAnimeRepository()
	jsonFilePath := "./anime.json"
	if err := repo.LoadAnimesFromJSON(jsonFilePath); err != nil {
		log.Fatalf("Failed to load anime data from JSON file: %v", err)
	}

	animeService := service.NewAnimeService(repo)

	animeController := controller.NewAnimeController(animeService)

	r := gin.Default()
	r.POST("/recommendations", animeController.GetRecommendations)
	r.POST("/evaluate/:id", animeController.EvaluateAnime)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
