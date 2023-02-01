package reverseProxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"gateway-app/config"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(c *gin.Context) {
	wordsService, err := url.Parse("http://" + config.WordsServiceAddress + config.WordsServicePort)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(wordsService)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = wordsService.Host
		req.URL.Scheme = wordsService.Scheme
		req.URL.Host = wordsService.Host
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
