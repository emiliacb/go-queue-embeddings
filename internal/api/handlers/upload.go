package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/emiliacb/go-queue-embeddings/internal/app/services"
)

func UploadHandler(c *gin.Context) {
	// TODO: Later we will use a file
	document := c.PostForm("text")

	enqueueService := services.NewEnqueueService()

	documentService, err := services.NewDocumentService()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id, err := enqueueService.Enqueue(document)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = documentService.SaveDocument(id, services.StatusProcessing)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	template := fmt.Sprintf(`
		<div hx-get="/documents/%s?format=htmx" hx-trigger="every 1s">loading...</div>
	`, id)

	c.String(200, template)
}
