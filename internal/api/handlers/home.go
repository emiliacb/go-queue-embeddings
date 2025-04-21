package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	html := `
	<html>
		<body>
			<h1>Go Queue Embeddings</h1>
			<p>This is a temporary home page for testing the huggingface space.</p>
		</body>
		</html>
	`
	c.Data(http.StatusOK, "text/html", []byte(html))
}
