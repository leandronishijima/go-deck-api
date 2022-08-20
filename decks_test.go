package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDeckWithoutCardsParameter(t *testing.T) {
	deck := NewDeck(false, nil)

	assert.False(t, deck.Shuffled, "Property shuffled needs to be follow the parameter")
	assert.NotNil(t, deck.DeckId, "DeckId should be not nil")
	assert.Len(t, deck.Cards, 52, "When not pass card parameters, expected all cards")
	assert.Equal(t, deck.Remaining, 52, "Should be initialize with deck size")
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

	assert.False(t, deck.Shuffled, "Property shuffled needs to be follow the parameter")
	assert.NotNil(t, deck.DeckId, "DeckId should be not nil")
	assert.Len(t, deck.Cards, 4, "Should be have all cards passed")
	assert.Equal(t, deck.Remaining, 4, "Should be initialize with deck size")
	assert.Equal(t, deck.Cards, cardsExpected, "Should be create all the cards from parameter")
}

func TestNewDeckWithShuffedParameter(t *testing.T) {
	deck := NewDeck(true, []string{"AS", "AC", "2C", "KH"})

	cardsExpected := []Card{
		{Value: "ACE", Suit: "SPADES", Code: "AS"},
		{Value: "ACE", Suit: "CLUBS", Code: "AC"},
		{Value: "2", Suit: "CLUBS", Code: "2C"},
		{Value: "KING", Suit: "HEARTS", Code: "KH"},
	}

	assert.NotNil(t, deck.DeckId, "DeckId should be not nil")
	assert.True(t, deck.Shuffled, "Property shuffled needs to be follow the parameter")
	assert.Len(t, deck.Cards, 4, "Should be have all cards passed")
	assert.Equal(t, deck.Remaining, 4, "Should be initialize with deck size")
	assert.NotEqual(t, deck.Cards, cardsExpected, "Should be create all the cards from parameter")
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
