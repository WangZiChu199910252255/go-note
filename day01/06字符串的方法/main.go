package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "wzc"
	s2 := "zdhqzc"
	s3 := "王子初"
	fmt.Printf("%T\n", s1)
	// 输出是s1的长度  			len
	fmt.Println(len(s1))
	// 字符串拼接				+或fmt.Sprintf
	fmt.Println(s1 + s2)
	ss := fmt.Sprintf("%s%s", s1, s2)
	fmt.Println(ss)
	// 字符串分割				strings.Split
	ret1 := strings.Split(ss, "z")
	fmt.Println(ret1)
	// 判断包含					strings.Contains
	ret2 := strings.Contains(ss, "w")
	fmt.Println(ret2)
	// 前缀 					strings.HasPrefix
	ret3 := strings.HasPrefix(ss, "w")
	fmt.Println(ret3)
	// 后缀						strings.HasSuffix
	ret4 := strings.HasSuffix(ss, "q")
	fmt.Println(ret4)
	// 判断字串第一次出现的位置   strings.Index
	ret5 := strings.Index(ss, "zc")
	fmt.Println(ret5)
	// 判断字串最后一次出现的位置 strings.LastIndex
	ret6 := strings.LastIndex(ss, "zc")
	fmt.Println(ret6)
	// 多个字符串拼接
	fmt.Println(strings.Join(ret1, "+"))
	// 修改字符串
	s4 := []rune(s3)
	s4[0] = '张'
	fmt.Println(string(s4))
	// 字符串和字符的区别
	c1 := "王" //string
	c2 := '王' //int32
	fmt.Printf("c1: %T c2: %T", c1, c2)

	str := []rune("王子初")
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}
}
