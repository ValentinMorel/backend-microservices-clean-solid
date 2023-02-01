package repositories

import (
	"math/rand"
	"strings"
	"words-app/models"
	"words-app/pkg/interfaces"

	"github.com/afex/hystrix-go/hystrix"
)

type WordRepositoryWithFeatures struct {
	WordRepository interfaces.IWordRepository
}

func (repository *WordRepositoryWithFeatures) AddOne(word string) (models.WordResponse, error) {
	output := make(chan models.WordResponse, 1)
	hystrix.ConfigureCommand("words/add", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("words/add", func() error {

		response, _ := repository.WordRepository.AddOne(word)
		output <- response
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.WordResponse{}, err
	}

}

func (repository *WordRepositoryWithFeatures) RandomSelect() (models.WordResponse, error) {
	output := make(chan models.WordResponse, 1)
	hystrix.ConfigureCommand("words/random_select", hystrix.CommandConfig{Timeout: 2000})
	errors := hystrix.Go("words/random_select", func() error {

		response, _ := repository.WordRepository.RandomSelect()

		output <- response

		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return models.WordResponse{}, err
	}
}

func (repository *WordRepositoryWithFeatures) GetAllWordsByPrefix(word string) (models.WordsResponse, error) {
	output := make(chan models.WordsResponse, 1)
	hystrix.ConfigureCommand("words/get_all_words_by_prefix", hystrix.CommandConfig{Timeout: 2000})
	errors := hystrix.Go("words/get_all_words_by_prefix", func() error {

		response, _ := repository.WordRepository.GetAllWordsByPrefix(word)

		output <- response
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return models.WordsResponse{}, err
	}

}

type WordRepository struct {
	WordRepository interfaces.IDbHandler
}

func (repository *WordRepository) GetAllWordsByPrefix(word string) (models.WordsResponse, error) {
	var response models.WordsResponse
	result, err := repository.WordRepository.GetAll()
	if err != nil {
		return models.WordsResponse{}, err
	}
	for _, v := range result {
		if strings.HasPrefix(v.Word, word) {
			response.Words = append(response.Words, v.Word)
		}
	}
	return response, nil
}
func (repository *WordRepository) RandomSelect() (models.WordResponse, error) {
	result, err := repository.WordRepository.GetAll()
	if err != nil {
		return models.WordResponse{}, err
	}
	upperBound := len(result)
	// Intn will fail if < 0
	if (upperBound-1)+1 > 0 {
		// id can't be 0 so set the minimum id to 1
		// -> (max - min + 1) + min
		id := rand.Intn((upperBound - 1 + 1) + 1)
		for _, v := range result {
			if v.Id == id {
				return v, nil
			}
		}
	}

	return models.WordResponse{}, nil
}
func (repository *WordRepository) AddOne(word string) (models.WordResponse, error) {

	return repository.WordRepository.AddOne(word)
}
