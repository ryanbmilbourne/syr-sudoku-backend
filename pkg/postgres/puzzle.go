package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	app "github.com/ryanbmilbourne/syr-sudoku-backend/pkg"
	log "github.com/sirupsen/logrus"
)

// A PostgreSQL implementation of app.PuzzleService interface
type PuzzleService struct {
	DB *sql.DB
}

func (s *PuzzleService) Init(dbUri string) error {
	if dbUri == "" {
		return errors.New("Empty database uri")
	}

	// Connect to the db
	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		return err
	}
	s.DB = db
	log.WithFields(log.Fields{
		"dbUri": dbUri,
	}).Debug("Connected to DB")

	// Create the table if it doesn't already exists.
	if _, err := s.DB.Exec(`CREATE TABLE IF NOT EXISTS puzzles (
                                   uuid serial primary key,
				   user_id text,
				   solution json,
				   create_timestamp TIMESTAMP DEFAULT now())`,
	); err != nil {
		return errors.Wrap(err, "Error creating database table")
	}

	return nil
}

func (s *PuzzleService) GetPuzzle(uuid string) (*app.Puzzle, error) {
	return nil, errors.New("Not yet implemented!")
}

func (s *PuzzleService) GetPuzzles() ([]*app.Puzzle, error) {
	return nil, errors.New("Not yet implemented!")
}

func (s *PuzzleService) CreatePuzzle(u *app.Puzzle) error {
	return errors.New("Not yet implemented!")
}

func (s *PuzzleService) DeletePuzzle(uuid string) error {
	return errors.New("Not yet implemented!")
}
