package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// START OMIT
var c chan string

func concGet(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return string(bodyBytes)
} // .concurentGet

func main() {
	c := make(chan string)

	go func() { c <- concGet("https://www.google.com") }()
	go func() { c <- concGet("https://www.yahoo.com") }()
	go func() { c <- concGet("https://www.cdc.gov") }()

	responses := []int{
		strings.Count(<-c, "div"), 
		strings.Count(<-c, "div"), 
		strings.Count(<-c, "div")}

	fmt.Println(responses)
} // .main
// END OMIT
