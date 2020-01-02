package main

import (
	".../handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/tokenAuth", handler.LoginHandler)

	r.Run(":8080")
}
