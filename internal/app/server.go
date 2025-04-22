package app

import (
	"github.com/emiliacb/go-queue-embeddings/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", handlers.HomeHandler)
	router.GET("/v2", handlers.HomeHandlerV2)
	router.Static("/static", "./static")
	router.GET("/health", handlers.HealthHandler)
	router.POST("/embed", handlers.EmbedHandler)
	router.POST("/documents", handlers.UploadHandler)
	router.GET("/documents/:id", handlers.DocumentHandler)
	return router
}

func StartServer() {
	router := setupRouter()
	router.Run(":8080")
}
