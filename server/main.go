package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleRoot)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Serving on PORT 8080...")
	}
}
