package main

import "fmt"

func main() {
	f1 := 123
	f2 := "wzc"
	// 查看变量的类型%T
	fmt.Printf("%T\n", f1)
	// 查看变量的值%v	加个#可以为类型加一个标识符
	fmt.Printf("%v\n", f1)
	fmt.Printf("%#v\n", f2)
	// 转化为二进制输出
	fmt.Printf("%b\n", f1)
	// 转化查看十进制输出
	fmt.Printf("%d\n", f1)
	// 转化为八进制输出
	fmt.Printf("%o\n", f1)
	// 转化为十六进制输出
	fmt.Printf("%x\n", f1)
	// 查看字符串
	fmt.Printf("%s\n", f2)
}
