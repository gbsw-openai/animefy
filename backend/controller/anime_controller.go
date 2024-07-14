package controllers

import (
	"context"
	"net/http"
	"strconv"

	models "backend/model"
	services "backend/service"

	"github.com/gin-gonic/gin"
)

type AnimeController struct {
	service *services.AnimeService
}

func NewAnimeController(service *services.AnimeService) *AnimeController {
	return &AnimeController{
		service: service,
	}
}

func (ctrl *AnimeController) GetRecommendations(c *gin.Context) {
	var userPreferences models.UserPreferences
	if err := c.ShouldBindJSON(&userPreferences); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recommendations, err := ctrl.service.RecommendAnime(context.Background(), userPreferences)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recommendations)
}

func (ctrl *AnimeController) EvaluateAnime(c *gin.Context) {
	animeIDStr := c.Param("id")
	animeID, err := strconv.Atoi(animeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anime ID"})
		return
	}

	var userPreferences models.UserPreferences
	if err := c.ShouldBindJSON(&userPreferences); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating, err := ctrl.service.EvaluateAnime(context.Background(), animeID, userPreferences)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rating": rating})
}
