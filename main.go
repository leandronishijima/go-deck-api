package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createDeckForm struct {
	Shuffled bool
	Cards    []string
}

var decks = []Deck{}

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/deck/new", createDeck)
		api.GET("/decks", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, decks) })
	}

	router.Run("localhost:8080")
}

func createDeck(c *gin.Context) {
	var req createDeckForm

	c.BindJSON(&req)
	new_deck := Deck{DeckId: uuid.New().String(), Suffled: req.Shuffled, Remaining: 52}

	decks = append(decks, new_deck)
	c.IndentedJSON(http.StatusOK, new_deck)
}
