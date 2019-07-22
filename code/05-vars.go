package main

import "fmt"

var c, python bool

const java = "jdk"

func main() {
	var (
		i int
		f float64
	)

	it := "hello"
	// var it string
	// it = "hello"

	fmt.Println(i, f, c, python, java)
	fmt.Printf("type: %T, val: %v\n", it, it)
}
