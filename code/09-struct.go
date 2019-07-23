package main

import (
	"fmt"
)

type container struct {
	a string
	b string
	c int
} // .type

func main() {
	c := container{
		"hello",
		"world",
		42,
	} // .c

	fmt.Printf("%+v", c)

} // .main
