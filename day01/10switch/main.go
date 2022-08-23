package main

import "fmt"

func main() {
	// 简化判断
	switch i := 1; i {
	case 0:
		fmt.Println("这是一个0")
	case 1:
		fmt.Println("这是一个1")
	case 2:
		fmt.Println("这是一个2")
	default:
		fmt.Println("不是数字")
	}
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("这是奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("这是偶数")
	default:
		fmt.Println(n)
	}
}
