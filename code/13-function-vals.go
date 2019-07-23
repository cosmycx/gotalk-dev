package main

import (
	"fmt"
)

type compFunc func(int, int) int

func compute(fn compFunc, a, b int) int {
	return fn(a, b)
} // .compute

func main() {

	add := func(x, y int) int {
		return x + y
	} // .add
	substr := func(x, y int) int {
		return x - y
	} // .substr

	fmt.Println(compute(add, 3, 4))
	fmt.Println(compute(substr, 3, 4))
} // .main
