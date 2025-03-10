package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/the-siili/BlackJackStt4/server/handlers"
	"github.com/the-siili/BlackJackStt4/server/middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.HandleRoot)

	loggedMux := middleware.Logging(mux)

	port := ":8080"
	fmt.Println("Server is running on http://localhost" + port)
	err := http.ListenAndServe(port, loggedMux)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
