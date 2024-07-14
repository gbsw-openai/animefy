package routes

import (
	controller "backend/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(animeController *controller.AnimeController) *gin.Engine {
	r := gin.Default()

	r.POST("/recommendations", animeController.GetRecommendations)
	r.POST("/evaluate/:id", animeController.EvaluateAnime)

	return r
}
