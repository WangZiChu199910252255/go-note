package main

import "fmt"

func test() {
	var ip *int
	fmt.Printf("ip: %v\n", ip) //nil
	fmt.Printf("ip: %T\n", ip) //*int
	var i int = 100
	ip = &i
	fmt.Printf("ip: %v\n", ip)  //地址 0x....
	fmt.Printf("ip: %v\n", *ip) //100
}

func test2() {
	a := []int{1, 2, 3, 4, 5}
	var i int
	ptr := make([]*int, len(a))
	for i = 0; i < len(a); i++ {
		ptr[i] = &a[i]
	}
	for i = 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
func main() {
	test2()
}
