// main.go
package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spiiderdan/design-patterns-in-go/creational/factory/handler"
	"github.com/spiiderdan/design-patterns-in-go/creational/factory/store"

)

func init() {
	_ = godotenv.Load()
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/shorten-url", handler.ShortenUrl)
	r.GET("/s/:shortUrl", handler.HandleShortUrlRedirect)

	store.ConnectToRedis(
		os.Getenv("REDIS_ADDRESS"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PASSWORD"),
		0,
	)

	r.Run(":" + os.Getenv("APP_PORT"))
}

