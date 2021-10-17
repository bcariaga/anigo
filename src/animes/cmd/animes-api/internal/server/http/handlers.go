package http

import (
	"net/http"

	"github.com/bcariaga/anigo/src/animes/internal/animes/fetching"
	"github.com/gin-gonic/gin"
)

func findAnimesHandlerBuilder(fetchService fetching.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		term := c.Query("search")
		animes, err := fetchService.FindByTerm(term)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": animes,
		})
	}
}
