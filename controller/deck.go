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
