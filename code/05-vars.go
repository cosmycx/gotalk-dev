package main

import "fmt"

// START OMIT
/*
bool, string
int  int8  int16  int32  int64, uint uint8 uint16 uint32 uint64 uintptr
byte (alias for uint8), rune ( alias for int32 // represents a Unicode code point)
float32 float64, complex64 complex128
*/
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
} // .main
// END OMIT
