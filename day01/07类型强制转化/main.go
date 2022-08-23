package main

import "fmt"

func main() {
	// 类型强制转化
	// 整形转浮点型
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Printf("f:%T  fValue:%v \n", f, f)
	// 字符串转rune和byte
	s1 := "王子初"
	sRune := []rune(s1)
	sBtye := []byte(s1)
	fmt.Printf("s1:%T sRune:%T sBtye:%T", s1, sRune, sBtye)
}
