package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"houkago-tea-time-api/internal/setup"
	"net/http"
)

type config struct {
	PORT string
	env  string
}

func main() {
	var cfg config

	// Add values for the config
	flag.StringVar(&cfg.PORT, "PORT", "4000", "API server port")
	flag.StringVar(&cfg.env, "mode", "test", "API current environment mode")
	flag.Parse()

	// Instantiate a router instance using the Gin engine with default options
	// and pass in a nil value to the Gin engine list of trusted proxies.
	r := gin.Default()
	_ = r.SetTrustedProxies(nil)
	gin.SetMode(cfg.env)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	setup.ConnectDatabase()

	err := r.Run(":" + cfg.PORT)
	if err != nil {
		return
	}
}
