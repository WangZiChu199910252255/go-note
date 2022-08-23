package main

import "fmt"

func main() {
	// type NewType int
	// var i NewType = 1
	// fmt.Printf("i: %T\n", i)
	type NewInt = int
	var i NewInt = 1
	fmt.Printf("i: %T\n", i)
}
