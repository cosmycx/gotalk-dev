package main

import (
	"fmt"
)

func main() {
	s1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // array
	s2 := []int{10, 11, 12, 13, 14}             // slice

	s2 = append(s1[:], s2...)

	for _, val := range s2 {

		if val%2 == 0 {
			fmt.Printf("number %v is even\n", val)
		} else {
			fmt.Printf("number %v is odd\n", val)
		}

	} //  .for

} // .main
