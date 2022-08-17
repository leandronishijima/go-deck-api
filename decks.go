package main

type Deck struct {
	DeckId    string `json:"deck_id"`
	Suffled   bool   `json:"suffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}
