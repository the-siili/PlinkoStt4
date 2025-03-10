package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/the-siili/BlackJackStt4/server/physics"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from server.")
}

func GenerateGame(w http.ResponseWriter, r *http.Request) {
	// Generate game data
	gameData, multiplier := physics.GenerateGame()
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Wrap the slice in a struct
	response := struct {
		Positions  [][2]float64 `json:"positions"`
		Multiplier float64      `json:"multiplier"`
	}{
		Positions:  gameData,
		Multiplier: multiplier,
	}

	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode and send JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
