package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createDeckForm struct {
	Shuffled bool
	Cards    []string
}

type drawDeckForm struct {
	Count int
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

func DrawCard(c *gin.Context) {
	var req drawDeckForm
	c.BindJSON(&req)
	deckId := c.Param("deck_id")

	var cardsDraw []Card
	for index, d := range decks {
		if d.DeckId == deckId {
			cardsDraw = d.DrawCard(req.Count)
			decks[index] = d
			break
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"cards": cardsDraw,
	})
}
