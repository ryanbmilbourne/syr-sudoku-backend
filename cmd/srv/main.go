package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ryanbmilbourne/syr-sudoku-backend/pkg/postgres"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	db *sql.DB
)

func GetPuzzle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ohai",
	})
}

func main() {
	log.SetLevel(log.InfoLevel)

	// Connect to DB
	dbUri := os.Getenv("DATABASE_URL")
	if dbUri == "" {
		log.Fatal("Need to specify DATABASE_URL env")
	}

	dbPuzzleService := postgres.PuzzleService{}

	if err := dbPuzzleService.Init(dbUri); err != nil {
		log.Fatalf("Could not connect to database: %q", err)
	}
	log.WithFields(log.Fields{
		"DATABASE_URL": dbUri,
	}).Info("Connected to database")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/puzzles", GetPuzzle)
	r.Run() // listen and serve on 0.0.0.0:8080
}
