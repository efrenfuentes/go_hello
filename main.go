package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type helloResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/static", staticHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := helloResponse{Message: "Hello World"}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	response := "The time is now " + time.Now().String()

	fmt.Fprintf(w, response)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}
