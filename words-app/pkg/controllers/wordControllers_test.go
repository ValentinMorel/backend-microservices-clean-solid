package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
	"words-app/models"
	"words-app/pkg/interfaces/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSearchByPrefix(t *testing.T) {

	wordService := new(mocks.IWordService)

	wordService.On("SearchByPrefix", "t").Return(models.WordsResponse{
		Words: []string{"trail", "task"},
	}, nil)

	wordController := WordController{wordService}

	req := httptest.NewRequest("GET", "http://localhost:8080/words/search?query=t", nil)
	w := httptest.NewRecorder()

	r := gin.New()
	r.GET("/words/search", wordController.SearchByPrefix)

	r.ServeHTTP(w, req)

	expectedResult := models.WordsResponse{
		Words: []string{"trail", "task"},
	}

	actualResult := models.WordsResponse{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
}
