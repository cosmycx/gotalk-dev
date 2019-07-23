package main

import "fmt"

type I interface {
	M()
}// .I

type T struct {
	S string
}// .T

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}// .M

func main() {
	var i I = T{"hello"}
	i.M()
}// .main