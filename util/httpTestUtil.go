package util

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/varadekd/card-game/helper"
)

func RequestAndDecodeResponse(method, api string, payload []byte, t *testing.T, router *gin.Engine) (helper.ResponseJSON, int) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, api, bytes.NewBuffer(payload))

	router.ServeHTTP(w, req)

	res := helper.ResponseJSON{}
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("We got and error %s while decoding response for %s", err.Error(), api)
	}

	return res, w.Code
}
