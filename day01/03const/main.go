package main

import "fmt"

const n = 100

// 批量声明常量时，如果某一行声明没有赋值，默认就和上面一样
const (
	a = 100
	b
)

// iota是go语言计算器，只能在常量的表达式中使用,类似枚举
const (
	a1 = iota //0
	a2        //1
	a3        //2
)
const (
	_  = iota
	KB = 1 << (10 * iota) //<<左移位数
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	const pi = 3.1415
	fmt.Println(pi)
	fmt.Println(n)
	fmt.Println(a, b)
	fmt.Println(a1, a2, a3)
}
