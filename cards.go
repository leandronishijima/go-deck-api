package main

import "strconv"

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

var orders = []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}
var suits_order = []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}

func NewCard(value, suit string) *Card {
	card := new(Card)
	card.Value = value
	card.Suit = suit
	card.setCode(value, suit)

	return card
}

func NewCardByCode(code string) *Card {
	card := new(Card)

	order, suit := getCardByCode(code)
	card.Value = order
	card.Suit = suit
	card.Code = code

	return card
}

func IsValidCode(code string) bool {
  order, suit := getCardByCode(code)

  var orderValid bool
  for _, o := range orders {
    if o == order {
      orderValid = true
    }
  }

  var suitValid bool
  for _, s := range suits_order {
    if s == suit {
      suitValid = true
    }
  }

  return orderValid && suitValid
}

func (card *Card) setCode(value, suit string) {
	if isNumber(value) {
		card.Code = value + suit[0:1]
	} else {
		card.Code = value[0:1] + suit[0:1]
	}
}

func isNumber(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func getCardByCode(code string) (string, string) {
	var order, suit string

	if len(code) == 3 {
		order = getOrderByCode(code[0:2])
	} else {
		order = getOrderByCode(code[0:1])
	}

	suit = getSuitByCode(code[len(code)-1:])

	return order, suit
}

func getOrderByCode(code string) string {
	switch code {
	case "A":
		return "ACE"

	case "J":
		return "JACK"

	case "Q":
		return "QUEEN"

	case "K":
		return "KING"

	default:
		return code
	}
}

func getSuitByCode(code string) string {
	switch code {
	case "S":
		return "SPADES"

	case "D":
		return "DIAMONDS"

	case "C":
		return "CLUBS"

	case "H":
		return "HEARTS"
	}

	return code
}
