package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/ryanbmilbourne/syr-sudoku-backend/pkg"
)

func GetPuzzle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ohai",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/puzzles", GetPuzzle)
	r.Run() // listen and serve on 0.0.0.0:8080
}
