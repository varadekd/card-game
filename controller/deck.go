package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/varadekd/card-game/helper"
	"github.com/varadekd/card-game/model"
	"golang.org/x/exp/slices"
)

var generatedDecks []model.Deck

func GeneratedDeck(c *gin.Context) {
	payload := model.GenerateDeckPayload{}
	response := helper.ResponseJSON{}

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		fmt.Errorf("Got an error while parsing new deck payload. Error: %s", err.Error())
		response.Success = false
		response.Error = "User shared and invalid payload"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	defaultDeck := []model.Card{}

	// Reading default deck from the file location
	filePath, _ := helper.GetEnvVariable("DEFAULT_CARDS_FILE_STORAGE")

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Errorf("Got an error '%s' while reading the default deck", err.Error())
		response.Success = false
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err = json.Unmarshal(data, &defaultDeck)

	if err != nil {
		fmt.Errorf("Got an error '%s' while un marshalling the default deck", err.Error())
		response.Success = false
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(defaultDeck) <= 0 {
		fmt.Errorf("Deck not found.")
		response.Success = false
		response.Error = "DeckID not found"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	newDeckID := uuid.New()

	deck := model.Deck{
		ID:      newDeckID,
		Shuffle: payload.Shuffle,
		GameID:  payload.GameID,
	}

	if len(payload.Cards) > 0 {
		for _, card := range defaultDeck {
			if slices.Contains(payload.Cards, card.Code) {
				deck.GeneratedDeck = append(deck.GeneratedDeck, card)
			}
		}
	} else {
		deck.GeneratedDeck = defaultDeck
	}

	// If shuffle is set to be true shuffling the generated cards using rand
	if payload.Shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(deck.GeneratedDeck), func(i, j int) {
			deck.GeneratedDeck[i], deck.GeneratedDeck[j] = deck.GeneratedDeck[j], deck.GeneratedDeck[i]
		})
	}

	// PlayingCards will have the value same as GeneratedCards since those cards are only been used by players.
	deck.PlayingCards = deck.GeneratedDeck
	deck.DeckSize = len(deck.GeneratedDeck)
	deck.CardsRemaining = len(deck.PlayingCards)
	deck.CreatedAt = time.Now()

	generatedDecks = append(generatedDecks, deck)

	response.Success = true
	response.Data = deck
	c.JSON(http.StatusCreated, response)
}

func OpenDeck(c *gin.Context) {

	deckID := c.Param("id")
	response := helper.ResponseJSON{}

	if deckID == "" {
		response.Error = "DeckID is missing"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err := uuid.Parse(deckID)

	if err != nil {
		response.Error = "DeckID is invalid"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	for _, deck := range generatedDecks {
		if deckID == deck.ID.String() {
			response.Success = true
			response.Data = deck
			c.JSON(http.StatusOK, response)
			return
		}
	}

	response.Error = "DeckID not found"
	c.JSON(http.StatusNotFound, response)
}

func DrawCardsFromDeck(c *gin.Context) {
	deckID := c.Param("id")
	response := helper.ResponseJSON{}

	if deckID == "" {
		response.Error = "DeckID is missing"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err := uuid.Parse(deckID)

	if err != nil {
		response.Error = "DeckID is invalid"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// payload verification
	payload := model.DrawCardFromDeckPayload{}

	err = c.ShouldBindJSON(&payload)

	if err != nil {
		fmt.Errorf("Got an error while parsing new deck payload. Error: %s", err.Error())
		response.Error = "User shared and invalid payload"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deckFound := false
	currentDeck := model.Deck{}
	currentDeckIndex := 0

	for index, deck := range generatedDecks {
		if deckID == deck.ID.String() {
			currentDeckIndex = index
			currentDeck = deck
			deckFound = true
			break
		}
	}

	if !deckFound {
		fmt.Errorf("We where unable to find deck with the deck ID: %s", deckID)
		response.Error = "DeckID not found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	if payload.CardsToBeDrawn > currentDeck.CardsRemaining {
		fmt.Errorf("Unable to draw cards from the deck, request %d cards but we only have %d remaining", payload.CardsToBeDrawn, currentDeck.CardsRemaining)
		response.Error = "There are no more cards left to be drawn from the deck"
		c.JSON(http.StatusConflict, response)
		return
	}

	drawnCards, remainingCards := drawCards(currentDeck.PlayingCards, payload.CardsToBeDrawn)

	// TODO: Add database for handling such updating of data, currently the data is
	// stored in the variable which will be vanished once the application is terminated.
	// Updating the current deck with remaining cards and updating count and time
	generatedDecks[currentDeckIndex].CardsRemaining = currentDeck.CardsRemaining - payload.CardsToBeDrawn
	generatedDecks[currentDeckIndex].PlayingCards = remainingCards
	generatedDecks[currentDeckIndex].DeckLastUsed = time.Now()

	response.Success = true
	response.Data = drawnCards
	c.JSON(http.StatusOK, response)
}

// drawCards will allow us to fetch cards from the deck.
// Currently this algorithm fetches the cards in array sequence.
// The function returns the drawn cards and also returns the remaining cards left in deck.
func drawCards(cards []model.Card, cardsToBeDrawn int) ([]model.Card, []model.Card) {
	rand.Seed(time.Now().UnixNano())

	drawnCards := make([]model.Card, cardsToBeDrawn)
	copy(drawnCards, cards[:cardsToBeDrawn])

	remainingCards := make([]model.Card, 0, len(cards)-cardsToBeDrawn)
	remainingCards = append(remainingCards, cards[cardsToBeDrawn:]...)

	return drawnCards, remainingCards
}
