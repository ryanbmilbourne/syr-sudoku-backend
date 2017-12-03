package syrSudokuBackend

type PuzzleState [9][9]uint8

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
