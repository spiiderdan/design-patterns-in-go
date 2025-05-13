package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spiiderdan/design-patterns-in-go/creational/factory/generator"
	"github.com/spiiderdan/design-patterns-in-go/creational/factory/store"
)

type URLShorteningRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	Method  string `json:"method"`
}

func ShortenUrl(c *gin.Context) {
	var req URLShorteningRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gen := generator.Factory(req.Method)
	short := gen.Generate(req.LongUrl)
	store.SaveMapping(short, req.LongUrl)

	host := "http://" + os.Getenv("APP_ADDRESS") + ":" + os.Getenv("APP_PORT") + "/s/"
	c.JSON(http.StatusOK, gin.H{
		"short_url": host + short,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	short := c.Param("shortUrl")
	long := store.RetrieveLongUrl(short)

	c.HTML(http.StatusOK, "redirect.html", gin.H{
		"LongURL": long,
	})
}