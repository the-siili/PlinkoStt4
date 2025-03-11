package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/the-siili/BlackJackStt4/server/physics"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from server.")
}

func GenerateGame(w http.ResponseWriter, r *http.Request) {
	// Generate game data
	gameData, multiplier := physics.GenerateGame()

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse JSON into struct
	payload := struct {
		Bet float64 `json:"bet"`
	}{}
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Check if bet is valid
	if payload.Bet < 0 {
		http.Error(w, "Invalid bet", http.StatusBadRequest)
	}
	win := payload.Bet * multiplier

	// TODO logic for saving bets and money etc.

	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Wrap the slice in a struct
	response := struct {
		Positions  [][2]float64 `json:"positions"`
		Multiplier float64      `json:"multiplier"`
		Win        float64      `json:"win"`
	}{
		Positions:  gameData,
		Multiplier: multiplier,
		Win:        win,
	}

	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode and send JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
