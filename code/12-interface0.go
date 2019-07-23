package main

import "fmt"

func main() {
	var a interface{}

	a = "hello"
	fmt.Println(a)
	a = 42
	fmt.Println(a)

	if b, ok := a.(int); ok {
		fmt.Println("b: ", b)
	}

	fmt.Println("a: ", a)
} // .main
