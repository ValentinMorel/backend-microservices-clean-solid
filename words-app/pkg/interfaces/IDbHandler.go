package interfaces

import (
	"words-app/models"
)

type IDbHandler interface {
	AddOne(word string) (models.WordResponse, error)
	GetAll() ([]models.WordResponse, error)
}
