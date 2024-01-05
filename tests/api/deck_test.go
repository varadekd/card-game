package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/varadekd/card-game/config"
	"github.com/varadekd/card-game/helper"
	"github.com/varadekd/card-game/util"
)

var router *gin.Engine

// We will temporarily store the deck ID from the TestGenerateDeck suite for further verification of other APIs.

var deckID string

func TestGenerateDeck(t *testing.T) {
	// Setting up router for the test execution
	router = config.SetupRouter()

	// Generating default card deck to ensure default deck is generated
	helper.GenerateDefaultDeck()

	t.Run("Generating default deck with shuffle false", func(t *testing.T) {

		payload := map[string]bool{
			"shuffle": false,
		}

		payloadString, err := json.Marshal(payload)

		if err != nil {
			t.Errorf("Test execution failed because of an error. Err: %s", err.Error())
		}

		res, code := util.RequestAndDecodeResponse("POST", "/deck/new", payloadString, t, router)

		// Verifying api status it should be 201
		assert.Equal(t, http.StatusCreated, code, fmt.Sprintf("We expected http status %d but got %d", http.StatusCreated, code))

		// Verify that the success field in the API response is set to 'true' to indicate success
		assert.Equal(t, true, res.Success, fmt.Sprintf("We expected the 'success' field to be set to true but found it as %t", res.Success))

		// Verifying that Data field is not empty, if it is empty failing the test case
		assert.NotNil(t, res.Data, fmt.Sprintf("We expected that the data within the response field should not be nil"))
	})

	t.Run("Generating custom deck with shuffle false", func(t *testing.T) {
		payload := map[string]any{
			"shuffle": false,
			"cards":   []string{"AD", "AC", "AH", "AS"},
		}

		payloadString, err := json.Marshal(payload)

		if err != nil {
			t.Errorf("Test execution failed because of an error. Err: %s", err.Error())
		}

		res, code := util.RequestAndDecodeResponse("POST", "/deck/new", payloadString, t, router)

		// Verifying api status it should be 201
		assert.Equal(t, http.StatusCreated, code, fmt.Sprintf("We expected http status %d but got %d", http.StatusCreated, code))

		// Verify that the success field in the API response is set to 'true' to indicate success
		assert.Equal(t, true, res.Success, fmt.Sprintf("We expected the 'success' field to be set to true but found it as %t", res.Success))

		// Verifying that Data field is not empty, if it is empty failing the test case
		assert.NotNil(t, res.Data, fmt.Sprintf("We expected that the data within the response field should not be nil"))

		if res.Data != nil {
			// For these api the data that will be returned will in form on map[string]interface{}
			resData := res.Data.(map[string]interface{})
			deckID = resData["_id"].(string)
		}

	})

	t.Run("Generating deck with invalid payload", func(t *testing.T) {
		payload := map[string]string{
			"shuffle": "false",
		}

		payloadString, err := json.Marshal(payload)

		if err != nil {
			t.Errorf("Test execution failed because of an error. Err: %s", err.Error())
		}

		res, code := util.RequestAndDecodeResponse("POST", "/deck/new", payloadString, t, router)

		// Verifying api status it should be 400
		assert.Equal(t, http.StatusBadRequest, code, fmt.Sprintf("We expected http status %d but got %d", http.StatusBadRequest, code))

		// Verify that the success field in the API response is set to 'false' to indicate failure
		assert.Equal(t, false, res.Success, fmt.Sprintf("We expected the 'success' field to be set to false but found it as %t", res.Success))

		// Verifying that Data field is empty, if it is not empty failing the test case
		assert.Nil(t, res.Data, fmt.Sprintf("We expected that the data within the response field should not nil"))

		// Verifying error message
		msg := "User shared and invalid payload"
		assert.Equal(t, msg, res.Error, fmt.Sprintf("We expected error message to be %s but found %s", msg, res.Error))
	})
}

func TestOpenDeck(t *testing.T) {

	t.Run("Fetching deck with valid deckID", func(t *testing.T) {
		api := fmt.Sprintf("/deck/%s", deckID)
		res, code := util.RequestAndDecodeResponse("GET", api, nil, t, router)

		// Verifying api status it should be 200
		assert.Equal(t, http.StatusOK, code, fmt.Sprintf("We expected http status %d but got %d", http.StatusOK, code))

		// Verify that the success field in the API response is set to 'true' to indicate success
		assert.Equal(t, true, res.Success, fmt.Sprintf("We expected the 'success' field to be set to true but found it as %t", res.Success))

		// Verifying that Data field is not empty, if it is empty failing the test case
		assert.NotNil(t, res.Data, fmt.Sprintf("We expected that the data within the response field should not be nil"))
	})

	t.Run("Fetching deck with invalid deckID", func(t *testing.T) {
		res, code := util.RequestAndDecodeResponse("GET", "/deck/invalidID", nil, t, router)

		// Verifying api status it should be 400
		assert.Equal(t, http.StatusBadRequest, code, fmt.Sprintf("We expected http status %d but got %d", http.StatusOK, code))

		// Verify that the success field in the API response is set to 'false' to indicate success
		assert.Equal(t, false, res.Success, fmt.Sprintf("We expected the 'success' field to be set to false but found it as %t", res.Success))

		// Verifying that Data field is not empty, if it is empty failing the test case
		assert.Nil(t, res.Data, fmt.Sprintf("We expected that the data within the response field should be nil"))

		// Verifying error message
		msg := "DeckID is invalid"
		assert.Equal(t, msg, res.Error, fmt.Sprintf("We expected error message to be %s but found %s", msg, res.Error))
	})

	t.Run("Fetching deck with deckID not part of data", func(t *testing.T) {
		newDeckID := uuid.New()

		api := fmt.Sprintf("/deck/%s", newDeckID.String())
		res, code := util.RequestAndDecodeResponse("GET", api, nil, t, router)

		// Verifying api status it should be 404
		assert.Equal(t, http.StatusNotFound, code, fmt.Sprintf("We expected http status %d but got %d", http.StatusNotFound, code))

		// Verify that the success field in the API response is set to 'false' to indicate success
		assert.Equal(t, false, res.Success, fmt.Sprintf("We expected the 'success' field to be set to false but found it as %t", res.Success))

		// Verifying that Data field is empty, if it is not empty failing the test case
		assert.Nil(t, res.Data, fmt.Sprintf("We expected that the data within the response field should be nil"))

		// Verifying error message
		msg := "DeckID not found"
		assert.Equal(t, msg, res.Error, fmt.Sprintf("We expected error message to be %s but found %s", msg, res.Error))
	})
}

func TestDrawCardsFromDeck(t *testing.T) {

	t.Run("Drawing valid numbers of cards that can be drawn from deck", func(t *testing.T) {
		payload := map[string]int{
			"cardsToBeDrawn": 2,
		}

		payloadString, err := json.Marshal(payload)

		if err != nil {
			t.Errorf("Test execution failed because of an error. Err: %s", err.Error())
		}

		api := fmt.Sprintf("/deck/%s/draw-cards", deckID)
		res, code := util.RequestAndDecodeResponse("PUT", api, payloadString, t, router)

		// Verifying api status it should be 200
		assert.Equal(t, http.StatusOK, code, fmt.Sprintf("We expected http status %d but got %d", http.StatusOK, code))

		// Verify that the success field in the API response is set to 'true' to indicate success
		assert.Equal(t, true, res.Success, fmt.Sprintf("We expected the 'success' field to be set to true but found it as %t", res.Success))

		// Verifying that Data field is not empty, if it is empty failing the test case
		assert.NotNil(t, res.Data, fmt.Sprintf("We expected that the data within the response field should not be nil"))
	})

	t.Run("Drawing invalid numbers of cards that can be drawn from deck", func(t *testing.T) {
		payload := map[string]int{
			"cardsToBeDrawn": 200,
		}

		payloadString, err := json.Marshal(payload)

		if err != nil {
			t.Errorf("Test execution failed because of an error. Err: %s", err.Error())
		}

		api := fmt.Sprintf("/deck/%s/draw-cards", deckID)
		res, code := util.RequestAndDecodeResponse("PUT", api, payloadString, t, router)

		// Verifying api status it should be 409
		assert.Equal(t, http.StatusConflict, code, fmt.Sprintf("We expected http status %d but got %d", http.StatusConflict, code))

		// Verify that the success field in the API response is set to 'false' to indicate success
		assert.Equal(t, false, res.Success, fmt.Sprintf("We expected the 'success' field to be set to false but found it as %t", res.Success))

		// Verifying that Data field is empty, if it is not empty failing the test case
		assert.Nil(t, res.Data, fmt.Sprintf("We expected that the data within the response field should not nil"))
	})

	t.Run("Drawing cards from the deckID not part of data", func(t *testing.T) {
		payload := map[string]int{
			"cardsToBeDrawn": 200,
		}

		payloadString, err := json.Marshal(payload)

		if err != nil {
			t.Errorf("Test execution failed because of an error. Err: %s", err.Error())
		}

		newDeckID := uuid.New()

		api := fmt.Sprintf("/deck/%s/draw-cards", newDeckID.String())

		res, code := util.RequestAndDecodeResponse("PUT", api, payloadString, t, router)

		// Verifying api status it should be 404
		assert.Equal(t, http.StatusNotFound, code, fmt.Sprintf("We expected http status %d but got %d", http.StatusNotFound, code))

		// Verify that the success field in the API response is set to 'false' to indicate success
		assert.Equal(t, false, res.Success, fmt.Sprintf("We expected the 'success' field to be set to false but found it as %t", res.Success))

		// Verifying that Data field is empty, if it is not empty failing the test case
		assert.Nil(t, res.Data, fmt.Sprintf("We expected that the data within the response field should not nil"))
	})

	t.Run("Drawing cards using invalid payload", func(t *testing.T) {
		payload := map[string]string{
			"cardsToBeDrawn": "10",
		}

		payloadString, err := json.Marshal(payload)

		if err != nil {
			t.Errorf("Test execution failed because of an error. Err: %s", err.Error())
		}

		api := fmt.Sprintf("/deck/%s/draw-cards", deckID)
		res, code := util.RequestAndDecodeResponse("PUT", api, payloadString, t, router)

		// Verifying api status it should be 400
		assert.Equal(t, http.StatusBadRequest, code, fmt.Sprintf("We expected http status %d but got %d", http.StatusCreated, code))

		// Verify that the success field in the API response is set to 'false' to indicate failure
		assert.Equal(t, false, res.Success, fmt.Sprintf("We expected the 'success' field to be set to false but found it as %t", res.Success))

		// Verifying that Data field is empty, if it is not empty failing the test case
		assert.Nil(t, res.Data, fmt.Sprintf("We expected that the data within the response field should not nil"))

		// Verifying error message
		msg := "User shared and invalid payload"
		assert.Equal(t, msg, res.Error, fmt.Sprintf("We expected error message to be %s but found %s", msg, res.Error))
	})

	t.Run("Drawing cards using invalid deckID", func(t *testing.T) {
		res, code := util.RequestAndDecodeResponse("PUT", "/deck/invalidID/draw-cards", nil, t, router)

		// Verifying api status it should be 400
		assert.Equal(t, http.StatusBadRequest, code, fmt.Sprintf("We expected http status %d but got %d", http.StatusOK, code))

		// Verify that the success field in the API response is set to 'false' to indicate success
		assert.Equal(t, false, res.Success, fmt.Sprintf("We expected the 'success' field to be set to false but found it as %t", res.Success))

		// Verifying that Data field is not empty, if it is empty failing the test case
		assert.Nil(t, res.Data, fmt.Sprintf("We expected that the data within the response field should be nil"))

		// Verifying error message
		msg := "DeckID is invalid"
		assert.Equal(t, msg, res.Error, fmt.Sprintf("We expected error message to be %s but found %s", msg, res.Error))
	})
}
