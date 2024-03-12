package moves

import (
	"ketchum/models"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewMoveRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

// GetAllMoves returns a list of all moves
func (r *Repo) GetAllMoves() ([]models.Move, error) {
	sql := `
	SELECT
	*
	FROM moves.moves m
	`
	moves := make([]models.Move, 0)
	err := r.db.Select(&moves, sql)
	if err != nil {
		return nil, err
	}
	return moves, nil
}

// GetMoveByName returns a move by its name
func (r *Repo) GetMoveByName(name string) (models.Move, error) {
	sql := `
	SELECT
	*
	FROM moves.moves m
	WHERE move_name = $1
	`
	move := models.Move{}
	err := r.db.Get(&move, sql, name)
	return move, err
}

// GetMoveByType returns a list of moves by type
func (r *Repo) GetMoveByType(moveType string) ([]models.Move, error) {
	sql := `
	SELECT
	*
	FROM moves.moves m
	WHERE move_type = $1
	`
	moves := make([]models.Move, 0)
	err := r.db.Select(&moves, sql, moveType)
	if err != nil {
		return nil, err
	}
	return moves, nil
}

// CreateMove creates a new move
func (r *Repo) CreateMove(move models.Move) error {
	sql := `
	INSERT INTO moves.moves
	(move_name, move_type, move_category, move_power, move_accuracy, move_pp, move_accuracy)
	VALUES
	($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(sql, move.Name, move.Type, move.Category, move.Power, move.Accuracy, move.PP, move.Accuracy)
	return err
}
