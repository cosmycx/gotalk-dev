package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", hello)
	log.Println("Starting server on port 4444")

	log.Fatal(http.ListenAndServe(":4444", nil))
} // .main

func hello(w http.ResponseWriter, r *http.Request) {

	log.Println("serving", r.URL)
	
	w.Write([]byte("Hello"))
} // .home
