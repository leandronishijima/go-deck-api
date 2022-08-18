package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCard(t *testing.T) {
	card := NewCard("ACE", "SPADES")

	assert.Equal(t, card.Value, "ACE", "Card value needs to be ACE")
	assert.Equal(t, card.Suit, "SPADES", "Card suit needs to be SPADES")
	assert.Equal(t, card.Code, "AS", "Card code needs to be the fist value from value and first letter of suit")
}

func TestGenerateCardsWithoutParams(t *testing.T) {
	assert.Equal(t, GenerateCards(), allCards(), "GenerateCards without parameters should return all cards")
}

func allCards() []Card {
	values := []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}
	suits := []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}

	cards := []Card{}
	for _, value := range values {
		for _, suit := range suits {
			cards = append(cards, *NewCard(value, suit))
		}
	}

	return cards
}
