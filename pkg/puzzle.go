package syrSudokuBackend

import "strconv"

type PuzzleState [9][9]uint8

func (p *PuzzleState) String() string {
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
	Solution PuzzleState `json:"solution,omitempty"`
}

type PuzzleService interface {
	GetPuzzle(uuid string) (*Puzzle, error)
	GetPuzzles() ([]*Puzzle, error)
	CreatePuzzle(u *Puzzle) error
	DeletePuzzle(uuid string) error
}
