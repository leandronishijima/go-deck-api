package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createDeckForm struct {
	Shuffled bool
	Cards    []string
}

func CreateDeck(c *gin.Context) {
	var req createDeckForm

	c.BindJSON(&req)
	newDeck, err := NewDeck(req.Shuffled, req.Cards)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		decks = append(decks, *newDeck)

		c.IndentedJSON(http.StatusOK, gin.H{
			"deck_id":   newDeck.DeckId,
			"shuffled":  newDeck.Shuffled,
			"remaining": newDeck.Remaining,
		})
	}
}

func OpenDeck(c *gin.Context) {
	deckId := c.Param("deck_id")

	for _, deck := range decks {
		if deck.DeckId == deckId {
			c.IndentedJSON(http.StatusOK, deck)
			return
		}
	}

	c.String(http.StatusNotFound, "Deck not found")
}
