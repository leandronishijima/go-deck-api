package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateCardsWithoutParams(t *testing.T) {
	assert.Equal(t, GenerateCards(nil), allCards(), "GenerateCards without parameters should return all cards")
}

func TestGenerateCardsWithParams(t *testing.T) {
	cards := []string{"AS", "AC", "2C", "KH"}

	expected := []Card{
		{Value: "ACE", Suit: "SPADES", Code: "AS"},
		{Value: "ACE", Suit: "CLUBS", Code: "AC"},
		{Value: "2", Suit: "CLUBS", Code: "2C"},
		{Value: "KING", Suit: "HEARTS", Code: "KH"},
	}

	assert.Equal(t, GenerateCards(cards), expected, "GenerateCards with parameters should return all cards declared")
}

func TestNewDeckWithoutCardsParameter(t *testing.T) {
	deck := NewDeck(false, nil)

	assert.False(t, deck.Suffled, "Property shuffled needs to be follow the parameter")
	assert.NotNil(t, deck.DeckId, "DeckId should be not nil")
	assert.Len(t, deck.Cards, 52, "When not pass card parameters, expected all cards")
	assert.Equal(t, deck.Cards, allCards(), "Should be create all the cards")
}

func TestNewDeckWithCardsParameter(t *testing.T) {
	deck := NewDeck(false, []string{"AS", "AC", "2C", "KH"})

	cardsExpected := []Card{
		{Value: "ACE", Suit: "SPADES", Code: "AS"},
		{Value: "ACE", Suit: "CLUBS", Code: "AC"},
		{Value: "2", Suit: "CLUBS", Code: "2C"},
		{Value: "KING", Suit: "HEARTS", Code: "KH"},
	}

	assert.False(t, deck.Suffled, "Property shuffled needs to be follow the parameter")
	assert.NotNil(t, deck.DeckId, "DeckId should be not nil")
	assert.Len(t, deck.Cards, 4, "When not pass card parameters, expected all cards")
	assert.Equal(t, deck.Cards, cardsExpected, "Should be create all the cards from parameter")
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
