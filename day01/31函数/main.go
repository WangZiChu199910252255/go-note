package main

import "fmt"

type add func(a int, b int) int

func test(a1 add) {
	fmt.Println(a1(1, 2))
}

func main() {
	func(a int) {
		fmt.Println(a)
	}(1)
	var a add = func(a int, b int) int {
		return a + b
	}
	res := a(1, 2)
	fmt.Println(res)
	test(a)
}
