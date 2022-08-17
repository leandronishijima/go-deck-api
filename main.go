package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type deck struct {
	DeckId    string `json:"deck_id"`
	Suffled   bool   `json:"suffled"`
	Remaining int    `json:"remaining"`
}

func createDeck(c *gin.Context) {
	new_deck := deck{DeckId: "a251071b-662f-44b6-ba11-e24863039c59", Suffled: false, Remaining: 52}
	c.IndentedJSON(http.StatusOK, new_deck)
}

func main() {
	router := gin.Default()

	router.POST("/api/deck/new", createDeck)

	// listen and serve on 0.0.0.0:8080
	router.Run()
}
