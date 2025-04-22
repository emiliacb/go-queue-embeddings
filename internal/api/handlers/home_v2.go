package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandlerV2(c *gin.Context) {
	c.HTML(http.StatusOK, "home_v2.html", gin.H{})
}
