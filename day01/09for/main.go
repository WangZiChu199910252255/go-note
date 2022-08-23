package main

import "fmt"

func main() {
	// for循环
	// 基本方式
	for i := 0; i < 10; i++ {
		// fmt.Println(i)
	}
	// 变种1
	i := 5
	for ; i < 10; i++ {
		// fmt.Println(i)
	}
	// 变种2
	t := 5
	for t < 10 {
		fmt.Println(t)
		t++
	}
	// 死循环
	// for {
	// 	fmt.Println("123")
	// }
	// for range键值循环
	s := "hello王子初"
	for i, v := range s {
		fmt.Printf("索引值：%d 键值：%c", i, v)
	}
	// 跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //当i等于五的时候跳出for循环
		}
	}
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //当I等于五的时候跳过这次循环
		}
		fmt.Println(i)
	}
	fmt.Println("OVER")
}
