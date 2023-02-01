package main

import (
	"words-app/config"
	"words-app/pkg/container"

	loggerMiddleware "words-app/tools"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	router.Use(gin.Recovery()) // to recover gin automatically
	router.Use(loggerMiddleware.JsonLoggerMiddleware())
	// Singleton on controller
	wordController := container.ServiceContainer().InjectController()

	// words service description
	wordRouter := router.Group("/words")
	wordRouter.POST("", wordController.AddOne)
	wordRouter.GET("/random", wordController.RandomSelect)
	wordRouter.GET("/search", wordController.SearchByPrefix)

	// run the service on the config defined port
	router.Run(config.WordsServicePort)
}
