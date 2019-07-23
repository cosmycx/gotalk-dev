package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(emoji("1F30E"))
} // .main

func emoji(hex string) (emj string, err error) {
	r, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		// log.Fatalf("Must exit error: %v", err)
		return emj, err
	} // .if
	return fmt.Sprintf("%c", r), nil
} // .main
