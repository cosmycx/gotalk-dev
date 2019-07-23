package main

import "fmt"

type Vertex struct {
	Lat, Long float64
} // .Vertex

var m map[string]Vertex

func main() {

	m = make(map[string]Vertex)

	m["Atlanta"] = Vertex{
		33.7, 84.3,
	} // .m
	
	fmt.Println(m["Atlanta"])
} // .main
