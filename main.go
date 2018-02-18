package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/", helloHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := helloResponse{Message: "Hello World"}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
