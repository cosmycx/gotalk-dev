package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Earth emoji: U+1F30E
	r, _ := strconv.ParseInt("1F30E", 16, 32)

	fmt.Printf("Hello %c\n", r)

} // .main
