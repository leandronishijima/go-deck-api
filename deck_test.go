package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateCardsWithoutParams(t *testing.T) {
	assert.Equal(t, GenerateCards(nil), allCards(), "GenerateCards without parameters should return all cards")
}

func TestGenerateCardsWithParams(t *testing.T) {
	cards := []string{"AS", "AC", "2C"}

	expected := []Card{
		{Value: "ACE", Suit: "SPADES", Code: "AS"},
		{Value: "ACE", Suit: "CLUBS", Code: "AC"},
		{Value: "2", Suit: "CLUBS", Code: "2C"},
	}

	assert.Equal(t, GenerateCards(cards), expected, "GenerateCards with parameters should return all cards declared")
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
