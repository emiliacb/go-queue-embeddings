package app

import (
	"github.com/gin-gonic/gin"
	"github.com/emiliacb/go-queue-embeddings/internal/api/handlers"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/health", handlers.HealthHandler)
	router.POST("/embed", handlers.EmbedHandler)
	return router
}

func StartServer() {
	router := setupRouter()
	router.Run(":8080")
}
