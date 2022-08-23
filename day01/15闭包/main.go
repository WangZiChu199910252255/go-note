package main

import "fmt"

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func cal(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(a int) int {
		base -= a
		return base
	}
	return add, sub
}
func main() {
	f1, f2 := cal(10)
	fmt.Printf("f1结果: %v f2结果: %v\n", f1(1), f2(2))
	fmt.Printf("f1结果: %v f2结果: %v\n", f1(3), f2(4))
	fmt.Printf("f1结果: %v f2结果: %v\n", f1(5), f2(6))
	fmt.Printf("f1结果: %v f2结果: %v\n", f1(7), f2(8))
	// f := add()
	// r := f(10)
	// fmt.Printf("r: %v\n", r)
	// r = f(20)
	// fmt.Printf("r: %v\n", r)
	// fmt.Printf("----------------------- \n")

	// f1 := add()
	// r = f1(100)
	// fmt.Printf("r: %v\n", r)
	// r = f1(200)
	// fmt.Printf("r: %v\n", r)
}
