package container

import (
	"sync"
	"words-app/models"
	"words-app/pkg/controllers"
	"words-app/pkg/db"
	"words-app/pkg/repositories"
	"words-app/pkg/services"
)

var (
	k             *ControllerBuilder
	containerOnce sync.Once
)

type IServiceContainer interface {
	InjectController() *controllers.WordController
}

type ControllerBuilder struct{}

func (k *ControllerBuilder) InjectController() *controllers.WordController {

	dbHandler := &db.DbHandler{
		Db: []models.WordResponse{},
	}

	wordRepository := &repositories.WordRepository{dbHandler}
	wordService := &services.WordService{&repositories.WordRepositoryWithFeatures{wordRepository}}
	wordController := &controllers.WordController{wordService}

	return wordController
}

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &ControllerBuilder{}
		})
	}
	return k
}
