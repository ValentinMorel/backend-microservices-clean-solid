package services

import (
	"testing"
	"words-app/models"
	"words-app/pkg/interfaces/mocks"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {

	wordRepository := new(mocks.IWordRepository)

	word1 := models.WordResponse{}
	word1.Id = 1
	word1.Word = "tail"

	word2 := models.WordResponse{}
	word2.Id = 2
	word2.Word = "task"

	wordRepository.On("AddOne", "tail").Return(word1, nil)

	wordService := WordService{wordRepository}

	expectedResult := models.WordResponse{
		Id:   1,
		Word: "tail",
	}

	actualResult, _ := wordService.AddOne("tail")
	assert.Equal(t, expectedResult, actualResult)

}

func TestSearchByPrefix(t *testing.T) {

	wordRepository := new(mocks.IWordRepository)

	word1 := models.WordResponse{}
	word1.Id = 1
	word1.Word = "tail"

	word2 := models.WordResponse{}
	word2.Id = 2
	word2.Word = "task"
	words := models.WordsResponse{
		Words: []string{word1.Word, word2.Word},
	}

	wordRepository.On("GetAllWordsByPrefix", "t").Return(words, nil)

	wordService := WordService{wordRepository}

	expectedResult := words

	actualResult, _ := wordService.SearchByPrefix("t")
	assert.Equal(t, expectedResult, actualResult)

}

func TestRandomSelect(t *testing.T) {

	wordRepository := new(mocks.IWordRepository)

	word1 := models.WordResponse{}
	word1.Id = 1
	word1.Word = "tail"

	word2 := models.WordResponse{}
	word2.Id = 2
	word2.Word = "task"

	wordRepository.On("RandomSelect").Return(word1, nil)

	wordService := WordService{wordRepository}

	expectedResult := word1

	actualResult, _ := wordService.RandomSelect()
	assert.Equal(t, expectedResult, actualResult)

}
