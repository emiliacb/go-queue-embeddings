package app

import (
	"github.com/emiliacb/go-queue-embeddings/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", handlers.HomeHandler)
	router.Static("/static", "./static")
	router.GET("/health", handlers.HealthHandler)
	router.POST("/embed", handlers.EmbedHandler)
	return router
}

func StartServer() {
	router := setupRouter()
	router.Run(":8080")
}
