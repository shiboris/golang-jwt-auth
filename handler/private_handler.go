package handler

import (
	"github.com/gin-gonic/gin"
)

// PrivateHandler :
func PrivateHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
	return
}
