package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	app "github.com/ryanbmilbourne/syr-sudoku-backend/pkg"
	"github.com/ryanbmilbourne/syr-sudoku-backend/pkg/grabber"
	"github.com/ryanbmilbourne/syr-sudoku-backend/pkg/postgres"

	"os"
	"time"

	log "github.com/sirupsen/logrus"
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
			"error": "Could not load image: " + err.Error(),
		})
		return
	}
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not load image: " + err.Error(),
		})
		return
	}

	// TODO: Do file header verification checks

	// Get dem Bytes
	var Buf bytes.Buffer
	var bytes []byte
	io.Copy(&Buf, src)
	bytes = Buf.Bytes()

	startParseTime := time.Now()
	puzzState, err := grabber.GrabPuzzle(bytes)
	totalParseTime := time.Since(startParseTime)
	log.WithFields(log.Fields{
		"timeElapsed": totalParseTime,
	}).Info("Parsed puzzle")
	puzzle.State = puzzState

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not preprocess image: " + err.Error(),
		})
		log.WithError(err)
		return
	}

	fmt.Printf(puzzState.String())

	// Solve that shit
	startSolveTime := time.Now()
	puzzSolution, err := grabber.SolvePuzzle(puzzState)
	totalSolveTime := time.Since(startSolveTime)
	log.WithFields(log.Fields{
		"timeElapsed": totalSolveTime,
	}).Info("Attmpted to solve puzzle")
	fmt.Printf(puzzSolution.String())

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not solve puzzle: " + err.Error(),
		})
		log.WithError(err)
		return
	}
	log.Info("Solved puzzle")

	puzzle.Solution = puzzSolution
	fmt.Printf(puzzSolution.String())

	// TODO: Save to Database

	c.JSON(201, puzzle)
}

func GetPuzzleHint(c *gin.Context) {
	var puzz app.Puzzle

	c.BindJSON(&puzz)

	fmt.Printf(puzz.State.String())

	startHintTime := time.Now()
	hintState, hintRow, hintCol, err := grabber.HintPuzzle(puzz.State)
	totalHintTime := time.Since(startHintTime)
	log.WithFields(log.Fields{
		"timeElapsed": totalHintTime,
	}).Info("Attmpted to fetch puzzle hint")

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not fetch hint: " + err.Error(),
		})
		return
	}
	log.Info("fetched hint")

	fmt.Println(hintState.String())

	c.JSON(200, gin.H{
		"state":      hintState,
		"hintCoords": [2]uint{hintRow, hintCol},
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
	r.POST("/puzzles/hints", GetPuzzleHint)
	r.POST("/puzzles", CreatePuzzle)
	r.Run() // listen and serve on 0.0.0.0:8080
}
