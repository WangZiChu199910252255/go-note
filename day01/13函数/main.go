package main

import (
	"fmt"
	_ "strconv"
)

func say() {
	fmt.Println("123")
}

func fn1(str string) {
	fmt.Println(str)
}

func fn2(name string, age uint) {
	fmt.Println(name, age)
}

func fn3(name string, age, age2 int) {
	fmt.Println(name, age, age2)
}

func fnAdd(num int) int {
	return num + 1
}

func fnAddTwo(num int) (int, int) {
	return num + 1, num + 2
}

func fnAddStr(num int) (numOne int, numTwo int) {
	numOne = num + 1
	numTwo = num + 2
	return
}

func sum(a, b int) int {
	return a + b
}

func max(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func sayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func test(name string, f func(string)) {
	f(name)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func cal(operator string) func(int, int) int {
	switch operator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func test2() {
	max := func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	r := max(1, 2)
	fmt.Printf("r: %v\n", r)
	r = func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}(1, 2)
	fmt.Printf("r: %v\n", r)
}

func test3() {
	name := "test"
	age := "100"
	f1 := func() string {
		return name + age
	}
	msg := f1()
	fmt.Printf("msg: %v\n", msg)
}
func main() {
	// test2()
	test3()
	// ff := cal("+")
	// r := ff(1, 2)
	// fmt.Printf("r:%v\n", r)
	// ff = cal("-")
	// r = ff(1, 2)
	// fmt.Printf("r:%v\n", r)
	// test("wzc", sayHello)
	// type f1 func(int, int) int
	// var ff f1
	// ff = sum
	// r := ff(1, 2)
	// fmt.Printf("r:%v\n", r)
	// ff = max
	// r = ff(1, 2)
	// fmt.Printf("r:%v", r)
	// say()
	// fn1("wzc")
	// fn2("wzc", 18)
	// fn3("wzc", 18, 19)
	// a := fnAdd(12)
	// fmt.Println(a)
	// a, b := fnAddTwo(12)
	// fmt.Println(a, b)
	// fmt.Println(fnAddStr(18))
	// _, c := fnAddStr(18)
	// fmt.Println(c)
}
