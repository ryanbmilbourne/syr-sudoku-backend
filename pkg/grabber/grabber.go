package grabber

import (
	opencv "github.com/go-opencv/go-opencv/opencv"
	app "github.com/ryanbmilbourne/syr-sudoku-backend/pkg"
)

func CreatePuzzle() (app.PuzzleState, error) {
	img := opencv.LoadImage("foo.jpg", 0)
	return nil, nil
}
