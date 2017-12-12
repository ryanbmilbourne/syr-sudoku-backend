package syrSudokuBackend

import "strconv"

type PuzzleState [][]uint

func NewPuzzleState() PuzzleState {
	a := make([][]uint, 9)
	for i := range a {
		a[i] = make([]uint, 9)
	}
	puzz := PuzzleState(a)
	return puzz
}

func (p PuzzleState) String() string {
	var out string
	for i := 0; i < len(p); i++ {
		if i == 3 || i == 6 {
			out = out + " -----------------------\n"
		}
		for j := 0; j < len(p[0]); j++ {
			if j == 3 || j == 6 {
				out = out + " | "
			}
			out = out + " " + strconv.Itoa(int(p[uint(i)][uint(j)]))
		}
		out = out + "\n"
	}
	return out
}

type Puzzle struct {
	UUID     string      `json:"puzzleId"`
	UserID   string      `json:"userId"`
	State    PuzzleState `json:"state" binding:"required"`
	Solution PuzzleState `json:"solution"`
}

type PuzzleService interface {
	GetPuzzle(uuid string) (*Puzzle, error)
	GetPuzzles() ([]*Puzzle, error)
	CreatePuzzle(u *Puzzle) error
	DeletePuzzle(uuid string) error
}
