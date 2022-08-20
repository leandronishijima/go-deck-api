package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var decks = []Deck{}

func setupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/deck/new", CreateDeck)
		api.GET("/deck/open/:deck_id", OpenDeck)
		api.GET("/decks", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, decks) })
	}

	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}
