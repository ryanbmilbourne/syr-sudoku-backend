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
	state := app.NewPuzzleState()

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

func SolvePuzzle(puzz app.PuzzleState) (app.PuzzleState, error, [2]uint) {
	solution, err := solve.SolvePuzzle(puzz)
	if err != nil {
		perr, _ := err.(*solve.PuzzleError)
		return app.PuzzleState{}, err, [2]uint{perr.Row, perr.Col}
	}

	return solution, nil, [2]uint{}
}

func HintPuzzle(puzz app.PuzzleState) (app.PuzzleState, uint, uint, error, [2]uint) {
	hintState, hintRow, hintCol, err := solve.Hint(puzz)
	if err != nil {
		perr, _ := err.(*solve.PuzzleError)
		return app.PuzzleState{}, 0, 0, err, [2]uint{perr.Row, perr.Col}
	}

	return hintState, hintRow, hintCol, nil, [2]uint{}
}
