package main

import "fmt"

func main() {
	// 整形
	// int
	// uint8 uint16 uint32 uint64 无符号整型
	// int8 int16 int32 int64 有符号整形
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%o\n", i1) //把十进制的数转化为八进制
	fmt.Printf("%x\n", i1) //把十进制的数转化为十六进制
	// 八进制   0开头
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制	0x开头
	i3 := 0xf
	fmt.Printf("%d\n", i3)
	// 查看变量类型
	fmt.Printf("%T\n", i3) //输出int
	i4 := int8(9)          //明确int的进制类型
	fmt.Printf("%o\n", i4)
	// 浮点数
	// float32 float64
	// math.MaxFloat32 //浮点数的最大值
	f1 := 1.2345
	fmt.Printf("%T\n", f1) //默认Go语言中的小数都是float64
	f2 := float32(1.2345)
	fmt.Printf("%T\n", f2) //显示float32类型
	// 复数
	var c1 complex64
	c1 = 1 + 2i
	fmt.Println(c1)
	// 布尔值
	// bool
	// 布尔值默认是false、不能强转
	var typeData bool
	fmt.Printf("%T value:%v\n", typeData, typeData)
	// 字符串
	// 内部使用的utf-8编码，字符串的值用双引号""来包裹
	// 不能使用''来包裹，go语言中''引号包裹的是字符(单独的字母、汉字、符号表示一个字符)
	s := 'w'
	s1 := "wzc"
	fmt.Println(s, s1)
	// 字符串转义符
	// \r 回车符、\n换行符\t制表符、\'单引号、\"双引号、\\反斜杠
	path := "\"D:\\GO\\src\\gitee.com\\wangzichu\""
	fmt.Println(path)
	// 多行字符串	用``来多行输出
	s2 := `
		王
		子
		初
	`
	fmt.Println(s2)
	const num = 1 / 2
	fmt.Println(num)
}
