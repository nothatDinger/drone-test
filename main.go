package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Header":  ctx.Request.Header,
			"message": "pong test",
		})
	})
	r.Run(":5000")
}
