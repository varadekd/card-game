package model

import (
	"time"

	"github.com/google/uuid"
)

type Deck struct {
	ID             uuid.UUID `json:"_id"`
	GameID         string    `json:"gameID"`
	Shuffle        bool      `json:"shuffle"`
	GeneratedDeck  []Card    `json:"generatedDeck"`
	PlayingCards   []Card    `json:"playingCards"`
	DeckSize       int       `json:"deckSize"`
	CardsRemaining int       `json:"cardRemaining"`
	DeckLastUsed   time.Time `json:"deckLastUsed"`
	CreatedAt      time.Time `json:"createdAt"`
}

// GenerateDeckPayload is used for creation on new deck
// GameID will be used to uniquely identify the deck used in that game.
// Shuffle true means the card sequence will be shuffled, false will be in sequence.
// Cards field will be used in case user wants only specific cards to be part of the deck
type GenerateDeckPayload struct {
	GameID  string   `json:"gameID"`
	Shuffle bool     `json:"shuffle"`
	Cards   []string `json:"cards"`
}
