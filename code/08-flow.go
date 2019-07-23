package main

import "fmt"

func main() {

	defer fmt.Println("first")

	for i := 0; i < 10; i++ {

		defer fmt.Println(i)

	} // .for

	fmt.Println("last")
} // .main

/*
for {
	break
	continue
}
*/
