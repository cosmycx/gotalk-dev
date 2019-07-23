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

// START OMIT
func getApi(apiPath string, c chan interface{}) {

	resp, _ := http.Get(apiPath)
	defer resp.Body.Close()
	var js interface{}
	json.NewDecoder(resp.Body).Decode(&js)

	c <- js
} // .getApi

func home(w http.ResponseWriter, r *http.Request) {

	c := make(chan interface{})
	
	go getApi("https://jsonplaceholder.typicode.com/todos/1", c)
	go getApi("https://jsonplaceholder.typicode.com/todos/2", c)
	go getApi("https://jsonplaceholder.typicode.com/todos/3", c)

	responses := []interface{}{<-c, <-c, <-c}
	json.NewEncoder(w).Encode(&responses)
} // .home
// END OMIT
