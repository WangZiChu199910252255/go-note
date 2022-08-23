package main

import "fmt"

// 单独声明
// var name string
// var age int
// var isOk bool
// 批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "wzc"
	age = 16
	isOk = true
	// var hehehe string
	// 在go语言中声明必须要调用
	fmt.Print(isOk)               //在终端中打印内容
	fmt.Printf("name:%s\n", name) //%s是占位符，使用name这个变量的值替换这个占位符
	fmt.Println(age)              //打印完之后会自动换行
	// hehehe = "嘿嘿嘿"

	// 声明变量同时赋值
	var s1 string = "wzc"
	fmt.Println(s1)
	// 类型推导（根据值来推断是什么类型）
	var s2 = 20
	fmt.Println(s2)
	// 简短变量声明只能在函数中使用
	s3 := "哈哈哈"
	s4, s5 := 19, "19"
	fmt.Println(s3, s4, s5)
	// 匿名变量声明 _ ,函数多参数时有些参数不需要使用可以使用_来声明
	// 同一个作用域中不能重复声明同名的变量
}
