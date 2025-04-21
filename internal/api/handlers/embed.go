package handlers

import (
	"net/http"

	"github.com/emiliacb/go-queue-embeddings/internal/app/domain"
	"github.com/emiliacb/go-queue-embeddings/internal/app/ports"
	"github.com/gin-gonic/gin"
)

// EmbedHandler is a temporary endpoint for development and debugging purposes.
// It provides direct access to the embedding functionality and will be removed
// once the proper /upload endpoint is implemented. This endpoint should not be
// used in production.
func EmbedHandler(c *gin.Context) {
	embedder := domain.GetContainer().Embedder

	text := c.PostForm("text")
	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "text parameter is required"})
		return
	}

	embedding, err := embedder.Embed(text, ports.EmbeddingConfig{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"embedding": embedding})
}
