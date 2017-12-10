package grabber

import (
	"errors"

	"github.com/jamesandersen/go-sudoku/sudokuparser"
	app "github.com/ryanbmilbourne/syr-sudoku-backend/pkg"
)

// GrabPuzzle parses a puzzle structure from a provided image
func GrabPuzzle(fileName string) (app.PuzzleState, error) {
	parsed, _ := sudokuparser.ParseSudokuFromFile(fileName)
	state := app.PuzzleState{}

	if parsed == "" {
		return state, errors.New("Unable to parse puzzle")
	}

	outerIdx := 0
	innerIdx := 0
	for i, val := range parsed {
		if (i != 0) && (i%9 == 0) {
			// New row!
			outerIdx++
			innerIdx = 0
		}

		if val == '.' {
			// Value is either unknown or empty
			state[outerIdx][innerIdx] = 0
		} else {
			// The cast here takes the ASCII value of the rune, so subtract
			// the value of ASCI 0 to get the true int value.
			state[outerIdx][innerIdx] = uint8(val - '0')
		}
		innerIdx++
	}
	return state, nil
}
