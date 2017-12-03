package syrSudokuBackend

type PuzzleState [9][9]uint8

type Puzzle struct {
	UUID     string
	UserID   string
	Solution PuzzleState
}

type PuzzleService interface {
	GetPuzzle(uuid string) (*Puzzle, error)
	GetPuzzles() ([]*Puzzle, error)
	CreatePuzzle(u *Puzzle) error
	DeletePuzzle(uuid string) error
}
