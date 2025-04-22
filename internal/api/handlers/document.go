package handlers

import (
	"fmt"

	"github.com/emiliacb/go-queue-embeddings/internal/app/services"
	"github.com/gin-gonic/gin"
)

func DocumentHandler(c *gin.Context) {
	id := c.Param("id")
	format := c.Query("format")

	documentService, err := services.NewDocumentService()
	if err != nil {
		fmt.Println("Error creating document service", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	status, err := documentService.ReadDocument(id)
	if err != nil {
		fmt.Println("Error reading document", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if format == "htmx" {
		c.String(200, "<div>Document %s is %s</div>", id, status)
		return
	}

	c.JSON(200, gin.H{
		"id":     id,
		"status": status.Status,
	})
}
