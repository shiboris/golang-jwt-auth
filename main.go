package main

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/golang-jwt-auth-test/handler"
)

func main() {
	r := gin.Default()

	r.POST("/tokenAuth", handler.LoginHandler)

	r.Run(":8080")
}
