package controller

import (
	"net/http"
	"strconv"

	repo "backend/repository"
	service "backend/service"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	service service.ChatService
}

func NewChatController(animeRepo repo.AnimeRepository, openaiKey string) *ChatController {
	return &ChatController{
		service: service.NewChatService(animeRepo, openaiKey),
	}
}

func (cc *ChatController) ChatWithBotHandler(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userID"))
	message := c.Query("message")

	response, err := cc.service.ChatWithBot(userID, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}
