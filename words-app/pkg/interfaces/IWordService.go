package interfaces

import "words-app/models"

type IWordService interface {
	AddOne(word string) (models.WordResponse, error)
	SearchByPrefix(prefix string) (models.WordsResponse, error)
	RandomSelect() (models.WordResponse, error)
}
