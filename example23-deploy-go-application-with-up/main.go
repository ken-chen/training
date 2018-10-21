package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":" + os.Getenv("PORT")
	stage := os.Getenv("UP_STAGE")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong " + stage,
		})
	})

	r.Run(port)
}
