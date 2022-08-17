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
