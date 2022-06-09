package repository

import "github.com/catehulu/rigged-coin/internal/models"

type DatabaseRepo interface {
	FindBoard(id string) *models.Board
	InsertBoard(board *models.Board)
	UpdateBoard(board *models.Board)
}
