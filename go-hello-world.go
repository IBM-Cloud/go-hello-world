package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/hi", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "hi",
		})
	})

	router.Run(":8080")
}
