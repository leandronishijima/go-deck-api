package main

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

var order = []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}
var suits_order = []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}

func NewCard(value, suit string) *Card {
	card := new(Card)
	card.Value = value
	card.Suit = suit
	card.Code = value[0:1] + suit[0:1]

	return card
}

func GenerateCards() []Card {
	cards := []Card{}
	for _, value := range order {
		for _, suit := range suits_order {
			cards = append(cards, *NewCard(value, suit))
		}
	}

	return cards
}
