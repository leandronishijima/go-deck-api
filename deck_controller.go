package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type createDeckForm struct {
	Shuffled bool
	Cards    []string
}

func CreateDeck(c *gin.Context) {
	var req createDeckForm

	c.BindJSON(&req)
	newDeck := NewDeck(req.Shuffled, nil)

	decks = append(decks, *newDeck)

	c.IndentedJSON(http.StatusOK, gin.H{
		"deck_id":   newDeck.DeckId,
		"shuffled":  newDeck.Shuffled,
		"remaining": newDeck.Remaining,
	})
}

func OpenDeck(c *gin.Context) {
	deckId := c.Param("deck_id")

	var deckFound Deck
	for _, deck := range decks {
		if deck.DeckId == deckId {
			deckFound = deck
			break
		}
	}

	c.IndentedJSON(http.StatusOK, deckFound)
}
