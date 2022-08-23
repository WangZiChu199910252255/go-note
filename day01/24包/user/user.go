package user

import "fmt"

// 大写开头的方法是共有的
func Say() {
	say()
	fmt.Println("Say")
}

// 小写只能在本包内使用
func say() {
	fmt.Println("say")
}
