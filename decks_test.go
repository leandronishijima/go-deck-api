package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDeckWithoutCardsParameter(t *testing.T) {
	deck, _ := NewDeck(false, nil)

	assert.False(t, deck.Shuffled, "Property shuffled needs to be follow the parameter")
	assert.NotNil(t, deck.DeckId, "DeckId should be not nil")
	assert.Len(t, deck.Cards, 52, "When not pass card parameters, expected all cards")
	assert.Equal(t, deck.Remaining, 52, "Should be initialize with deck size")
	assert.Equal(t, deck.Cards, allCards(), "Should be create all the cards")
}

func TestNewDeckWithCardsParameter(t *testing.T) {
	deck, _ := NewDeck(false, []string{"AS", "AC", "2C", "KH"})

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
	deck, _ := NewDeck(true, []string{"AS", "AC", "2C", "KH"})

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

func TestNewDeckWithInvalidCardCode(t *testing.T) {
	deck, err := NewDeck(false, []string{"AAS", "AC", "2C", "KH"})

	assert.Nil(t, deck)
	assert.Equal(t, "All the cards must be valid", err.Error())
}

func TestDraw(t *testing.T) {
	deck, _ := NewDeck(false, []string{"AS", "KD", "AC", "2C", "KH"})

	cardsExpected := []Card{
		{Value: "ACE", Suit: "SPADES", Code: "AS"},
		{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
		{Value: "ACE", Suit: "CLUBS", Code: "AC"},
		{Value: "2", Suit: "CLUBS", Code: "2C"},
		{Value: "KING", Suit: "HEARTS", Code: "KH"},
	}
	assert.Equal(t, cardsExpected, deck.Cards)
	assert.Equal(t, 5, deck.Remaining)

	drawCardExpected := []Card{{Value: "ACE", Suit: "SPADES", Code: "AS"}}
	drawCards, _ := deck.DrawCard(1)
	assert.Equal(t, drawCardExpected, drawCards)

	cardsExpectedAfterDraw := []Card{
		{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
		{Value: "ACE", Suit: "CLUBS", Code: "AC"},
		{Value: "2", Suit: "CLUBS", Code: "2C"},
		{Value: "KING", Suit: "HEARTS", Code: "KH"},
	}

	assert.Equal(t, cardsExpectedAfterDraw, deck.Cards)
	assert.Equal(t, 4, deck.Remaining)
}

func TestDrawCountAboveTheNumberOfCardsInDeck(t *testing.T) {
	deck, _ := NewDeck(false, []string{"AS", "KD"})

	_, err := deck.DrawCard(3)

	assert.Equal(t, "Number invalid of cards to draw, available: 2", err.Error())
}

func TestDrawWhenTheDeckIsEmpty(t *testing.T) {
	deck, _ := NewDeck(false, []string{"AS"})
	deck.DrawCard(1)

	_, err := deck.DrawCard(1)

	assert.Equal(t, "The deck is empty", err.Error())
}

func allCards() []Card {
	values := []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}
	suits := []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}

	cards := []Card{}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, *NewCard(value, suit))
		}
	}

	return cards
}
