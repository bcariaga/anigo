package http

import (
	"net/http"

	"github.com/bcariaga/anigo/src/animes/internal/animes/fetching"

	"github.com/gin-gonic/gin"
)

func MainHandler(
	fetchingService fetching.Service,
) (http.Handler, error) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/animes", findAnimesHandlerBuilder(fetchingService))
	return r, nil
}
