package main

import "github.com/google/uuid"

type Deck struct {
	DeckId    string `json:"deck_id"`
	Suffled   bool   `json:"suffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

func NewDeck(suffled bool, cards []string) *Deck {
	deck := new(Deck)
	deck.DeckId = uuid.NewString()
	deck.Suffled = suffled
	deck.Cards = GenerateCards(cards)

	return deck
}

func GenerateCards(cards []string) []Card {
	if cards == nil {
		return generateFullDeck()
	} else {
		return generateDeck(cards)
	}
}

func generateFullDeck() []Card {
	cards := []Card{}
	for _, value := range order {
		for _, suit := range suits_order {
			cards = append(cards, *NewCard(value, suit))
		}
	}

	return cards
}

func generateDeck(cards []string) []Card {
	deck := []Card{}
	for _, cardCode := range cards {
		deck = append(deck, *NewCardByCode(cardCode))
	}

	return deck
}
