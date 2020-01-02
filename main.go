package main

import (
	"./handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	u := r.Group("/auth")
	{
		u.GET("", handler.)
		u.POST("", handler.)
	}

	r.Run(":8080")
}
