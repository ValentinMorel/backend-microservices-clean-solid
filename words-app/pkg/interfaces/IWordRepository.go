package interfaces

import "words-app/models"

type IWordRepository interface {
	AddOne(word string) (models.WordResponse, error)
	GetAllWordsByPrefix(word string) (models.WordsResponse, error)
	RandomSelect() (models.WordResponse, error)
	//GetAll() ([]models.WordResponse, error)
}
