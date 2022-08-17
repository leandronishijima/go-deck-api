package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type deck struct {
	DeckId    string `json:"deck_id"`
	Suffled   bool   `json:"suffled"`
	Remaining int    `json:"remaining"`
}

var decks = []deck{}

func main() {
	router := gin.Default()

	router.POST("/api/deck/new", createDeck)
	router.GET("/api/decks", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, decks) })

	router.Run("localhost: 8080")
}

func createDeck(c *gin.Context) {
	new_deck := deck{DeckId: uuid.New().String(), Suffled: false, Remaining: 52}

	decks = append(decks, new_deck)
	c.IndentedJSON(http.StatusOK, new_deck)
}
