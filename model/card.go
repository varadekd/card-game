package model

// The Card struct assigns values, codes, and suits to cards:
// 1. The Value field indicates the card's value.
// 2. The Code field provides a unique identifier for the card within the deck.
// 3. The Suit field specifies the suit to which this card belongs.

type Card struct {
	Value string `json:"value"`
	Code  string `json:"code:"`
	Suit  string `json:"suit"`
}
