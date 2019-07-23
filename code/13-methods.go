package main

import (
	"fmt"
)

type container struct {
	a, b string
} // .type

func (c container) list() {
	fmt.Println(c.a, c.b)
} // list

func main() {
	c := container{
		"hello",
		"world",
	} // .c

	c.list()

} // .main
