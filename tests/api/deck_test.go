package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/varadekd/card-game/config"
	"github.com/varadekd/card-game/helper"
	"github.com/varadekd/card-game/util"
)

var router *gin.Engine

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
		assert.Equal(t, http.StatusBadRequest, code, fmt.Sprintf("We expected http status %d but got %d", http.StatusCreated, code))

		// Verify that the success field in the API response is set to 'false' to indicate failure
		assert.Equal(t, false, res.Success, fmt.Sprintf("We expected the 'success' field to be set to false but found it as %t", res.Success))

		// Verifying that Data field is not empty, if it is empty failing the test case
		assert.NotNil(t, res.Data, fmt.Sprintf("We expected that the data within the response field should not be nil"))

		// Verifying error message
		msg := "User shared and invalid payload"
		assert.Equal(t, msg, res.Error, fmt.Sprintf("We expected error message to be %s but found %s", msg, res.Error))
	})
}
