package main

import "fmt"

func main() {
	// if条件判断
	age := 18
	if age > 18 {
		fmt.Println("你成年了")
	} else {
		fmt.Println("你没成年")
	}
	// 多条件判断
	if age > 35 {
		fmt.Println("该退休了")
	} else if age > 18 {
		fmt.Println("好好学习")
	} else {
		fmt.Println("毕业了")
	}
	// 特殊写法
	if ageTwo := 19; ageTwo > 18 {
		fmt.Println("你成年了", ageTwo)
	} else {
		fmt.Println("你没成年", ageTwo)
	}
}
