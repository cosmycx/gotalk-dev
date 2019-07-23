package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	log.Println("Starting server on port 4444")
	log.Fatal(http.ListenAndServe(":4444", nil))
} // .main

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("serving", r.URL)

	type myjson struct {
		A, B string
		C    int
	} // .myjson
	json.NewEncoder(w).Encode(myjson{"hello", "world", 42})
} // .home
