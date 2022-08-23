package main

import "fmt"

func main() {
	var (
		a        = 5
		b        = 2
		boolTest = true
	)
	//算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句，不能写在等号的右边赋值 不能写 a = a--
	b--
	fmt.Println(a, b)
	//关系运算符
	fmt.Println(a == b) //判断a，b是否相同强类型同种类型才能比较
	fmt.Println(a != b) //不等于
	fmt.Println(a >= b) //大于等于
	fmt.Println(a <= b) //小于等于
	fmt.Println(a > b)  //大于
	fmt.Println(a < b)  //小于
	//逻辑运算符
	fmt.Println(a > 2 && a < 6) //a大于2且a小于6
	fmt.Println(a > 2 || a < 6) //a大于2或a小于6
	fmt.Println(!boolTest)      //boolTest取反变成false
	// 位运算符（二进制运算）
	var i1 = 5                //101
	var i2 = 2                //010
	fmt.Printf("%b\n", i1&i2) //&按位与 ，参与运算的两个数按位相与（两位均为1才为1）
	fmt.Printf("%b\n", i1|i2) //|按位或，参与运算的两个数按位或运算（两位有一个为1，就是1）
	fmt.Printf("%b\n", i1^i2) //^按位异或，参数运算的两个数按位异或运算（两位不一样则为1）
	fmt.Printf("%b\n", i1<<i2) //<<将二进制左移指定位数
	fmt.Printf("%b\n", i1>>i2) //>>将二进制右移指定位数，做位运算的时候不能超过变量的值
	// 赋值运算符(给变量赋值)
	var i3 int
	i3 = 10
	i3 += 1 //i3 = i3 + 1
	i3 -= 1	//i3 = i3 - 1
	i3 *= 2 //i3 = i3 * 2
	i3 /= 2 //i3 = i3 / 2
	i3 %= 2 //i3 = i3 % 2
	i3 <<= 2 //i3 = i3 << 2
	i3 &= 2 //i3 = i3 & 2
	i3 |= 2 //i3 = i3 | 2
}
