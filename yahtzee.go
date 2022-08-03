package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	username = "admin"
	password = "snakeeyes"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Player struct {
	Name string `json:"name"`
}

type Yahtzee struct {
	playerName string
	dice       map[int]*Die
}

func sendResponse(w http.ResponseWriter, responseCode int, status string, payload interface{}) {
	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	response := Response{
		Status: status,
		Data:   payload,
	}
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error occurred in JSON marshal: %s", err)
	}
	w.Write(jsonResp)
}

// Extract the die ID from the request
func parseRequestId(w http.ResponseWriter, r *http.Request) int {
	// Check ID variable exists
	vars := mux.Vars(r)
	if vars == nil {
		sendResponse(w, http.StatusBadRequest, "failed", "Die ID not found in request")
		return -1
	}
	// Try to get ID from request
	rawId, ok := vars["id"]
	if !ok {
		sendResponse(w, http.StatusBadRequest, "failed", "Die ID not found in request")
		return -1
	}
	// Parse ID to an integer
	id, err := strconv.Atoi(rawId)
	if err != nil || id < 1 || id > 5 {
		sendResponse(w, http.StatusBadRequest, "failed", "Die ID must be an integer between 1 and 5")
		return -1
	}
	return id
}

// Validate username and password
func checkAuthorization(w http.ResponseWriter, r *http.Request) bool {
	// Get basic auth
	suppliedUsername, suppliedPassword, ok := r.BasicAuth()
	// Check basic auth
	if !ok {
		sendResponse(w, http.StatusUnauthorized, "failed", "Basic auth missing or could not be parsed")
		return false
	}
	// Verify username and password are correct
	if suppliedUsername != username || suppliedPassword != password {
		sendResponse(w, http.StatusUnauthorized, "failed", "Incorrect username or password provided")
		return false
	}
	return true
}

// Initialize
func (y *Yahtzee) Init() {
	// Set default player name
	y.playerName = "SAS"
	// Initialize dice map
	y.dice = make(map[int]*Die)
	// Create 5 dice
	for id := 1; id <= 5; id++ {
		// Create die
		d := &Die{
			Id:    id,
			Value: 1,
		}
		// Roll die to initialize it
		d.Roll()
		// Add die to map
		y.dice[id] = d
	}
}

// Check if all dice show the same value
func (y *Yahtzee) isYahtzee(w http.ResponseWriter, r *http.Request) {
	// Determine if all dice are the same
	allEqual := true
	v := y.dice[1].Value
	for i := 2; i <= 5; i++ {
		if y.dice[i].Value != v {
			allEqual = false
			break
		}
	}
	// Send response
	sendResponse(w, http.StatusOK, "success", allEqual)
}

// Returns the name of the player
func (y *Yahtzee) getPlayerName(w http.ResponseWriter, r *http.Request) {
	// Send response
	sendResponse(w, http.StatusOK, "success", y.playerName)
}

// Sets the name of the player, requires authorization
func (y *Yahtzee) setPlayerName(w http.ResponseWriter, r *http.Request) {
	// Check authorization
	if !checkAuthorization(w, r) {
		return
	}
	// Parse request body
	var p Player
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&p)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, "failed", "Invalid request body format")
		return
	}
	// Set player name
	y.playerName = p.Name
	// Send empty success response message
	w.WriteHeader(http.StatusNoContent)
}

// Roll all five dice
func (y *Yahtzee) rollDice(w http.ResponseWriter, r *http.Request) {
	// Roll each die and make a list of them to return
	dice := make([]*Die, 5)
	for i := 0; i < 5; i++ {
		y.dice[i+1].Roll()
		dice[i] = y.dice[i+1]
	}
	// Send response
	sendResponse(w, http.StatusOK, "success", dice)
}

// Roll an individual die
func (y *Yahtzee) rollDie(w http.ResponseWriter, r *http.Request) {
	// Parse ID from request
	id := parseRequestId(w, r)
	if id == -1 {
		return
	}
	// Roll die
	y.dice[id].Roll()
	// Return die
	sendResponse(w, http.StatusOK, "success", y.dice[id])
}

// Returns the values of all five dice
func (y *Yahtzee) getDice(w http.ResponseWriter, r *http.Request) {
	// Transform dice from a map into a list
	dice := make([]*Die, 5)
	for i := 0; i < 5; i++ {
		dice[i] = y.dice[i+1]
	}
	// Send response
	sendResponse(w, http.StatusOK, "success", dice)
}

// Return the value of a single die
// Supports various media types for different representations of die value
func (y *Yahtzee) getDie(w http.ResponseWriter, r *http.Request) {
	// Parse ID from request
	id := parseRequestId(w, r)
	if id == -1 {
		return
	}
	// Return die in the proper format based on Accept header
	switch r.Header.Get("Accept") {
	case "":
		fallthrough
	case "*/*":
		fallthrough
	case "application/json":
		fallthrough
	case "application/vnd.yahtzee.int+json":
		sendResponse(w, http.StatusOK, "success", y.dice[id])
	case "application/vnd.yahtzee.float+json":
		sendResponse(w, http.StatusOK, "success", y.dice[id].AsFloat())
	case "application/vnd.yahtzee.word+json":
		sendResponse(w, http.StatusOK, "success", y.dice[id].AsWord())
	case "application/vnd.yahtzee.dots+json":
		sendResponse(w, http.StatusOK, "success", y.dice[id].AsDots())
	default:
		sendResponse(w, http.StatusBadRequest, "failed", "Invalid value in 'Accept' header")
	}
}

// Set the value of an individual die, requires authorization
func (y *Yahtzee) setDie(w http.ResponseWriter, r *http.Request) {
	// Check authorization
	if !checkAuthorization(w, r) {
		return
	}
	// Parse request body
	var d Die
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&d)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, "failed", "Invalid request body format")
		return
	}
	// Verify die id is valid (between 1 and 5)
	if d.Id < 1 || d.Id > 5 {
		sendResponse(w, http.StatusBadRequest, "failed", "Die ID must be between 1 and 5")
		return
	}
	// Verify die value is valid (between 1 and 6)
	if d.Value < 1 || d.Value > 6 {
		sendResponse(w, http.StatusBadRequest, "failed", "Die value must be between 1 and 6")
		return
	}
	// Update die value
	y.dice[d.Id].Value = d.Value
	// Send empty success response
	w.WriteHeader(http.StatusNoContent)
}
