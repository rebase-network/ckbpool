package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	startServer()
}

func startServer() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "ckbpool",
		})
	})

	fmt.Println("HTTP Server on: http://localhost:7000/")
	r.Run(":7000")
}
