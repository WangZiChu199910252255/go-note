package main

import "fmt"

func test(n int) int {
	if n == 1 {
		return 1
	}
	return n * test(n-1)
}

func main() {
	r := test(6)
	fmt.Println(r)
}
