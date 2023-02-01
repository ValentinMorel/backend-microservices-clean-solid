package main

import (
	"gateway-app/config"
	reverseProxyWords "gateway-app/reverse-proxy"
	loggerMiddleware "gateway-app/tools"

	"github.com/gin-gonic/gin"
)

func main() {
	service := gin.New()

	service.Use(gin.Recovery()) // to recover gin automatically
	service.Use(loggerMiddleware.JsonLoggerMiddleware())

	// Group by service to subrouters
	// words service description
	wordService := service.Group("/words")
	wordService.POST("", reverseProxyWords.ReverseProxy)
	wordService.GET("/random", reverseProxyWords.ReverseProxy)
	wordService.GET("/search", reverseProxyWords.ReverseProxy)

	// Run the gateway that will reach services
	service.Run(config.ProxyServicePort)
}
