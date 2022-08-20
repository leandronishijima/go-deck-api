package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var decks = []Deck{}

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/deck/new", CreateDeck)
		api.GET("/deck/open/:deck_id", OpenDeck)
		api.GET("/decks", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, decks) })
	}

	router.Run("localhost:8080")
}
