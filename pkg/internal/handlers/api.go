package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is coming from an API call",
	})
}
