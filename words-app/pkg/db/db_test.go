package db

import (
	"testing"
	"words-app/models"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {

	db := DbHandler{
		Db: []models.WordResponse{},
	}
	actualResult, err := db.AddOne("test")

	expectedResult := models.WordResponse{
		Id:   1,
		Word: "test",
	}
	assert.Equal(t, expectedResult, actualResult)
	assert.Nil(t, err)

}

func TestGetAll(t *testing.T) {

	db := DbHandler{
		Db: []models.WordResponse{},
	}
	_, _ = db.AddOne("test")
	actualResult, err := db.GetAll()
	expectedResult := []models.WordResponse{
		models.WordResponse{
			Id:   1,
			Word: "test",
		},
	}
	assert.Equal(t, expectedResult, actualResult)
	assert.Nil(t, err)
}
