package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateDeck(t *testing.T) {
	response := callApi(http.MethodPost, "/api/deck/new", `{"shuffled": false}`)

	assert.Equal(t, http.StatusOK, response.Code)

	bodyRes := convertBodyToMap(response.Body.String())

	assert.NotNil(t, bodyRes["deck_id"])
	assert.Equal(t, false, bodyRes["shuffled"])
	assert.Equal(t, float64(52), bodyRes["remaining"])
}

func TestCreateDeckWithShuffledTrueAndCards(t *testing.T) {
	response := callApi(
		http.MethodPost,
		"/api/deck/new",
		`{"shuffled": true, "cards": ["AS", "KD", "AC", "2C", "KH"]}`)

	assert.Equal(t, http.StatusOK, response.Code)

	bodyRes := convertBodyToMap(response.Body.String())

	assert.NotNil(t, bodyRes["deck_id"])
	assert.Equal(t, true, bodyRes["shuffled"])
	assert.Equal(t, float64(5), bodyRes["remaining"])
}

func TestCreateDeckWithInvalidCards(t *testing.T) {
	response := callApi(
		http.MethodPost,
		"/api/deck/new",
		`{"shuffled": false, "cards": ["NOK"]}`)

	assert.Equal(t, http.StatusBadRequest, response.Code)

	bodyRes := convertBodyToMap(response.Body.String())

	assert.Equal(t, "All the cards must be valid", bodyRes["error"])
}

func TestOpenDeck(t *testing.T) {
	response := callApi(
		http.MethodPost,
		"/api/deck/new",
		`{"shuffled": false, "cards": ["AS", "KD"]}`)

	bodyRes := convertBodyToMap(response.Body.String())
	deckId := bodyRes["deck_id"]

	responseGet := callApi(
		http.MethodGet,
		fmt.Sprintf("/api/deck/open/%s", deckId),
		"")

	bodyResGet := convertBodyToMap(responseGet.Body.String())

	assert.Equal(t, deckId, bodyResGet["deck_id"])
	assert.Equal(t, false, bodyResGet["shuffled"])
	assert.Equal(t, float64(2), bodyResGet["remaining"])

	cards := []interface{}([]interface{}{
		map[string]interface{}{"code": "AS", "suit": "SPADES", "value": "ACE"},
		map[string]interface{}{"code": "KD", "suit": "DIAMONDS", "value": "KING"},
	})

	assert.Equal(t, cards, bodyResGet["cards"])
}

func TestOpenDeckWhenDeckIdDoesntExist(t *testing.T) {
	callApi(
		http.MethodPost,
		"/api/deck/new",
		`{"shuffled": false, "cards": ["AS", "KD"]}`)

	responseGet := callApi(
		http.MethodGet,
		"/api/deck/open/1-wrong-not-uuid",
		"")

	assert.Equal(t, http.StatusNotFound, responseGet.Code)
}

func TestDrawCard(t *testing.T) {
	responseCreate := callApi(
		http.MethodPost,
		"/api/deck/new",
		`{"shuffled": false, "cards": ["AS", "KD", "QH"]}`)

	bodyRes := convertBodyToMap(responseCreate.Body.String())
	deckId := bodyRes["deck_id"]

	responseDraw := callApi(
		http.MethodPatch,
		fmt.Sprintf("/api/deck/%s/draw", deckId),
		`{"count": 1}`,
	)

	bodyResDraw := convertBodyToMap(responseDraw.Body.String())

	assert.Equal(t, http.StatusOK, responseDraw.Code)

	cards := []interface{}([]interface{}{
		map[string]interface{}{"code": "AS", "suit": "SPADES", "value": "ACE"},
	})

	assert.Equal(t, cards, bodyResDraw["cards"])
}

func TestDrawAllCards(t *testing.T) {
	responseCreate := callApi(
		http.MethodPost,
		"/api/deck/new",
		`{"shuffled": false, "cards": ["AS", "KD", "QH"]}`)

	bodyRes := convertBodyToMap(responseCreate.Body.String())
	deckId := bodyRes["deck_id"]

	responseDraw := callApi(
		http.MethodPatch,
		fmt.Sprintf("/api/deck/%s/draw", deckId),
		`{"count": 3}`,
	)

	bodyResDraw := convertBodyToMap(responseDraw.Body.String())

	assert.Equal(t, http.StatusOK, responseDraw.Code)

	cards := []interface{}([]interface{}{
		map[string]interface{}{"code": "AS", "suit": "SPADES", "value": "ACE"},
		map[string]interface{}{"code": "KD", "suit": "DIAMONDS", "value": "KING"},
		map[string]interface{}{"code": "QH", "suit": "HEARTS", "value": "QUEEN"},
	})

	assert.Equal(t, cards, bodyResDraw["cards"])
}

func TestDrawCardSequence(t *testing.T) {
	responseCreate := callApi(
		http.MethodPost,
		"/api/deck/new",
		`{"shuffled": false, "cards": ["AS", "KD"]}`)

	bodyRes := convertBodyToMap(responseCreate.Body.String())
	deckId := bodyRes["deck_id"]

	_ = callApi(
		http.MethodPatch,
		fmt.Sprintf("/api/deck/%s/draw", deckId),
		`{"count": 1}`,
	)

	_ = callApi(
		http.MethodPatch,
		fmt.Sprintf("/api/deck/%s/draw", deckId),
		`{"count": 1}`,
	)

	responseGet := callApi(
		http.MethodGet,
		fmt.Sprintf("/api/deck/open/%s", deckId),
		"")

	bodyResGet := convertBodyToMap(responseGet.Body.String())

	assert.Equal(t, deckId, bodyResGet["deck_id"])
	assert.Equal(t, float64(0), bodyResGet["remaining"])

	assert.Equal(t, []interface{}([]interface{}{}), bodyResGet["cards"])
}

func TestDrawCardsAboveTheCapacity(t *testing.T) {
	responseCreate := callApi(
		http.MethodPost,
		"/api/deck/new",
		`{"shuffled": false, "cards": ["AS"]}`)

	bodyRes := convertBodyToMap(responseCreate.Body.String())
	deckId := bodyRes["deck_id"]

	responseDraw := callApi(
		http.MethodPatch,
		fmt.Sprintf("/api/deck/%s/draw", deckId),
		`{"count": 2}`,
	)

	bodyResDraw := convertBodyToMap(responseDraw.Body.String())

	assert.Equal(t, http.StatusUnprocessableEntity, responseDraw.Code)
	assert.Equal(t, "Number invalid of cards to draw, available: 1", bodyResDraw["error"])
}

func TestTryToDrawCardWithoutCountParam(t *testing.T) {
	response := callApi(
		http.MethodPatch,
		"/api/deck/1/draw",
		"",
	)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, "{\"error\":\"Parameter 'count' is required\"}", response.Body.String())
}

func TestTryToDrawCardsFromEmptyDeck(t *testing.T) {
	responseCreate := callApi(
		http.MethodPost,
		"/api/deck/new",
		`{"shuffled": false, "cards": ["AS"]}`)

	bodyRes := convertBodyToMap(responseCreate.Body.String())
	deckId := bodyRes["deck_id"]

	callApi(
		http.MethodPatch,
		fmt.Sprintf("/api/deck/%s/draw", deckId),
		`{"count": 1}`,
	)

	response := callApi(
		http.MethodPatch,
		fmt.Sprintf("/api/deck/%s/draw", deckId),
		`{"count": 1}`,
	)

	bodyResError := convertBodyToMap(response.Body.String())

	assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	assert.Equal(t, "The deck is empty", bodyResError["error"])
}

func setupTestApi() (*gin.Engine, *httptest.ResponseRecorder) {
	router := setupRouter()
	w := httptest.NewRecorder()

	return router, w
}

func makeBodyReq(param string) *bytes.Reader {
	return bytes.NewReader([]byte(param))
}

func callApi(method, uri, body string) *httptest.ResponseRecorder {
	router, responseRecorder := setupTestApi()

	bodyReader := makeBodyReq(body)
	req, _ := http.NewRequest(method, uri, bodyReader)

	router.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func convertBodyToMap(bodyString string) map[string]interface{} {
	bodyRes := map[string]interface{}{}
	json.Unmarshal([]byte(bodyString), &bodyRes)

	return bodyRes
}
