package config_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varadekd/card-game/config"
)

func TestPing(t *testing.T) {
	router := config.SetupRouter()
	t.Run("Testing ping on server", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ping", nil)

		router.ServeHTTP(w, req)

		var data map[string]string

		if err := json.NewDecoder(w.Body).Decode(&data); err != nil {
			t.Fatalf("Failed to decode JSON: %v", err)
		}

		assert.EqualValues(t, http.StatusOK, w.Code, fmt.Sprintf("We expected http status %d but got %d", http.StatusOK, w.Code))
		assert.Equal(t, "pong", data["message"], fmt.Sprintf("We expected the message 'pong', but received %s.", data["message"]))

	})
}
