package main

import (
	"github.com/gin-gonic/gin"
	"houkago-tea-time-api/internal/setup"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	setup.ConnectDatabase()

	err := r.Run()
	if err != nil {
		return
	}
}
