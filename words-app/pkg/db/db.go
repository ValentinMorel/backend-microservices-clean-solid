package db

import (
	"errors"
	"words-app/models"
)

type DbHandler struct {
	Db []models.WordResponse
}

func (handler *DbHandler) AddOne(word string) (models.WordResponse, error) {

	var startLen int = 0
	var lastId int = 0
	if len(handler.Db) > 0 {
		lastId = handler.Db[len(handler.Db)-1].Id
	}
	var id int = lastId + 1
	itemToInsert := models.WordResponse{
		Id:   id,
		Word: word,
	}
	handler.Db = append(handler.Db, itemToInsert)
	// Check that the item is well inserted by increasing number of element by 1
	if startLen < len(handler.Db) {
		return itemToInsert, nil
	} else {
		return models.WordResponse{}, errors.New("Couldn't insert new word to database")
	}

}

func (handler *DbHandler) GetAll() ([]models.WordResponse, error) {
	return handler.Db, nil
}
