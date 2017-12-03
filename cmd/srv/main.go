package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	app "github.com/ryanbmilbourne/syr-sudoku-backend/pkg"
	"github.com/ryanbmilbourne/syr-sudoku-backend/pkg/postgres"

	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	db *sql.DB
)

func GetPuzzle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ohai",
	})
}

func CreatePuzzle(c *gin.Context) {
	puzzle := app.Puzzle{}

	puzzle.UserID = c.PostForm("userId")
	if puzzle.UserID == "" {
		c.JSON(400, gin.H{
			"error": "Missing `userId` in POST body",
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could load load image:" + err.Error(),
		})
		return
	}
	file.Filename = "/tmp/" + puzzle.UserID + "_" + time.Now().Format(time.RFC3339)
	c.SaveUploadedFile(file, file.Filename)
	log.WithFields(log.Fields{
		"Filename": file.Filename,
	}).Info("Rx file upload")

	c.JSON(201, puzzle)
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
	r.POST("/puzzles", CreatePuzzle)
	r.Run() // listen and serve on 0.0.0.0:8080
}
