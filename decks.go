package main

import (
	"errors"
	"math/rand"

	"github.com/google/uuid"
)

type Deck struct {
	DeckId    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

func NewDeck(shuffled bool, cards []string) (*Deck, error) {
	if cards != nil {
		for _, codeCard := range cards {
			if !IsValidCode(codeCard) {
				return nil, errors.New("All the cards must be valid")
			}
		}
	}

	deck := new(Deck)
	deck.DeckId = uuid.NewString()
	deck.Shuffled = shuffled
	deck.Cards = generateCards(cards)
	deck.Remaining = len(deck.Cards)

	if shuffled {
		shuffleDeck(deck.Cards)
	}

	return deck, nil
}

func (deck *Deck) DrawCard(count int) []Card {
	cards := deck.Cards[0:count]

	deck.Cards = deck.Cards[count:len(deck.Cards)]
	deck.Remaining = len(deck.Cards)

	return cards
}

func generateCards(cards []string) []Card {
	if cards == nil {
		return generateFullDeck()
	} else {
		return generateDeck(cards)
	}
}

func shuffleDeck(cards []Card) {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

func generateFullDeck() []Card {
	cards := []Card{}
	for _, suit := range suits_order {
		for _, value := range orders {
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
