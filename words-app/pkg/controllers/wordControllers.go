package controllers

import (
	"log"
	"net/http"
	"words-app/models"

	"words-app/pkg/interfaces"

	"github.com/gin-gonic/gin"
)

type WordController struct {
	WordService interfaces.IWordService
}

func (controller *WordController) RandomSelect(context *gin.Context) {

	response, err := controller.WordService.RandomSelect()
	if err != nil {
		context.JSON(http.StatusRequestTimeout, nil)
	}

	context.JSON(http.StatusOK, response)
}

func (controller *WordController) AddOne(context *gin.Context) {
	var request models.WordRequest
	context.BindJSON(&request)
	if len(request.Word) > 255 {
		context.JSON(http.StatusRequestEntityTooLarge, nil)
	} else {
		response, err := controller.WordService.AddOne(request.Word)
		if err != nil {
			log.Println("error: ", err)
		}

		context.JSON(http.StatusOK, response)
	}

}

func (controller *WordController) SearchByPrefix(context *gin.Context) {
	value, _ := context.GetQuery("query")
	response, err := controller.WordService.SearchByPrefix(value)
	if err != nil {
		context.JSON(http.StatusRequestTimeout, nil)
	}

	context.JSON(http.StatusOK, response)
}
