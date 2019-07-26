package main

import (
	"fmt"

	"github.com/gookit/color"

	"github.com/gin-gonic/gin"
)

var (
	putf  = fmt.Printf
	putln = fmt.Println
	puts  = fmt.Sprintf

	red   = color.FgRed.Render
	green = color.FgGreen.Render
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

	putf(green("\nHTTP Server on: http://localhost:7000/\n"))

	r.Run(":7000")
}
