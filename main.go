package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/htanmo/shorturl/handler"
	"github.com/htanmo/shorturl/store"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", handler.CreateShortUrl)

	r.GET("/:shortUrl", handler.HandleShortUrlRedirect)

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
