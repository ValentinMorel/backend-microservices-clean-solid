package services

import (
	"words-app/models"
	"words-app/pkg/interfaces"
)

type WordService struct {
	WordRepository interfaces.IWordRepository
}

func (service *WordService) AddOne(word string) (models.WordResponse, error) {
	response, err := service.WordRepository.AddOne(word)
	if err != nil {
		return models.WordResponse{}, err
	}
	return response, nil
}

func (service *WordService) RandomSelect() (models.WordResponse, error) {
	response, err := service.WordRepository.RandomSelect()
	if err != nil {
		return models.WordResponse{}, err
	}
	return response, nil
}

func (service *WordService) SearchByPrefix(word string) (models.WordsResponse, error) {
	response, err := service.WordRepository.GetAllWordsByPrefix(word)
	if err != nil {
		return models.WordsResponse{}, err
	}
	return response, nil
}
