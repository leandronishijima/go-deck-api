package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCard(t *testing.T) {
	tableTest := []map[string]string{
		{"value": "ACE", "suit": "SPADES", "code": "AS"},
		{"value": "2", "suit": "SPADES", "code": "2S"},
		{"value": "9", "suit": "SPADES", "code": "9S"},
		{"value": "JACK", "suit": "CLUBS", "code": "JC"},
		{"value": "QUEEN", "suit": "DIAMONDS", "code": "QD"},
		{"value": "10", "suit": "SPADES", "code": "10S"},
	}

	for _, test := range tableTest {
		card := NewCard(test["value"], test["suit"])

		assert.Equal(t, card.Value, test["value"], fmt.Sprintf("Card value should be %s", test["value"]))
		assert.Equal(t, card.Suit, test["suit"], fmt.Sprintf("Card suit should be %s", test["suit"]))
		assert.Equal(t, card.Code, test["code"], "Card code should be the value and first letter of suit")
	}
}

func TestNewCardByCode(t *testing.T) {
	tableTest := []map[string]string{
		{"value": "ACE", "suit": "SPADES", "code": "AS"},
		{"value": "2", "suit": "SPADES", "code": "2S"},
		{"value": "9", "suit": "SPADES", "code": "9S"},
		{"value": "JACK", "suit": "CLUBS", "code": "JC"},
		{"value": "QUEEN", "suit": "DIAMONDS", "code": "QD"},
		{"value": "10", "suit": "SPADES", "code": "10S"},
		{"value": "KING", "suit": "HEARTS", "code": "KH"},
	}

	for _, test := range tableTest {
		card := NewCardByCode(test["code"])

		assert.Equal(t, card.Value, test["value"], fmt.Sprintf("Card value should be %s", test["value"]))
		assert.Equal(t, card.Suit, test["suit"], fmt.Sprintf("Card suit should be %s", test["suit"]))
		assert.Equal(t, card.Code, test["code"], "Card code should be the value and first letter of suit")
	}
}
