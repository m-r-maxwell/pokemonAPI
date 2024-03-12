package moves

import "ketchum/models"

type IMoveRepo interface {
	GetMoveByName(name string) (models.Move, error)
	GetMoveByType(moveType string) ([]models.Move, error)
	GetAllMoves() ([]models.Move, error)
	CreateMove(move models.Move) error
}
