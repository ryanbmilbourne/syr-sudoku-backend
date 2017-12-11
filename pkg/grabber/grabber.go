package grabber

import (
	"errors"

	"github.com/Wrenky/sudoKu/solve"
	app "github.com/ryanbmilbourne/syr-sudoku-backend/pkg"
	"github.com/ryanbmilbourne/syr-sudoku-backend/pkg/sudokuparser"
)

// GrabPuzzle parses a puzzle structure from a provided image
func GrabPuzzle(bytes []byte) (app.PuzzleState, error) {
	parsed, _ := sudokuparser.ParseSudokuFromByteArray(bytes)
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
			state[outerIdx][innerIdx] = uint(val - '0')
		}
		innerIdx++
	}
	return state, nil
}

func SolvePuzzle(puzz app.PuzzleState) (app.PuzzleState, error) {
	solution, err := solve.SolvePuzzle(puzz)
	if err != nil {
		return app.PuzzleState{}, err
	}

	return solution, nil
}
