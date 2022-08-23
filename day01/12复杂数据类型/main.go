package main

import "fmt"

func test() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	s2 := s1[0:3]
	fmt.Printf("s2: %v\n", s2)
	s3 := s1[3:]
	fmt.Printf("s3: %v\n", s3)
	s4 := s1[:]
	fmt.Printf("s4: %v\n", s4)
	for i := 0; i < len(s1); i++ {
		fmt.Printf("s1[%v]:%v\n", i, s1[i])
	}
}

func test2() {
	var a1 = [...]int{1, 2, 3, 4, 5, 6}
	a2 := a1[:]
	fmt.Printf("a1: %v\n", a2)
}

func test3() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	for i, v := range s1 {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}
}

func test4() {
	var s1 = []int{}
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	fmt.Printf("s1: %v\n", s1)
}

func test5() {
	var s1 = []int{1, 2, 3}
	s1 = append(s1[:1], s1[2:]...)
	fmt.Printf("s1: %v\n", s1)
}

func test6() {
	var s1 = []int{1, 2, 3}
	s1[1] = 100
	fmt.Printf("s1: %v\n", s1)
}

func test7() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	var key = 5
	for i, v := range s1 {
		if v == key {
			fmt.Printf("i: %v v: %v", i, v)
		}
	}
}

func test8() {
	var s1 = []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copy(s2, s1)
	fmt.Printf("s1: %v s2: %v", s1, s2)
}

func test9() {
	var s1 = []int{1, 2}
	s1 = append(s1, 3)
	// fmt.Printf("s1: %v\n", s1)
	p1 := &s1[1]
	p2 := &s1[2]
	fmt.Printf("p1: %v %v\n", p1, *p1)
	fmt.Printf("p2: %v %v\n", p2, *p2)
}
func main() {
	test9()
	// test()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	// test8()
	// 数组
	// 如何定义数组
	// 定义数组需要指定存放数据的类型和长度 数组的长度是数组类型的一部分
	// var a1 [3]int
	// var a2 [4]int
	// // a1和a2不能做比较
	// fmt.Printf("a1:%T a2:%T\n", a1, a2)
	// // 数组的初始化
	// // 如果初始化，默认元素都是0值
	// fmt.Println(a1, a2)
	// // 1.初始化方式1
	// a1 = [3]int{1, 2, 3}
	// fmt.Println(a1)
	// // 2.初始化方式2
	// // 根据初始值自动推断数组的长度是多少
	// a10 := [...]int{1, 2, 3, 4, 5, 6}
	// fmt.Printf("a10:%T\n", a10)
	// // 3.初始化方式3
	// // 根据索引来初始化
	// a3 := [5]int{0: 1, 4: 2}
	// fmt.Println(a3)
	// // 数组的遍历
	// citys := [...]string{"北京", "上海", "深圳"}
	// // 1.根据索引来遍历
	// for i := 0; i < len(citys); i++ {
	// 	fmt.Println(citys[i])
	// }
	// // 2.for range遍历
	// for i, v := range citys {
	// 	fmt.Println(i, v)
	// }
	// // 多维数组
	// // [[1 2] [3 4] [5 6]]
	// a11 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	// fmt.Println(a11)
	// //多维数组的遍历
	// for i, v := range a11 {
	// 	fmt.Println(i, v)
	// 	for j, l := range v {
	// 		fmt.Println(j, l)
	// 	}
	// }
	// // 数组是值类型
	// b1 := [3]int{1, 2, 3}
	// b2 := b1
	// b2[0] = 100
	// fmt.Println(b1, b2)
	// // 数组练习
	// mun := 0
	// arr := [...]int{1, 3, 5, 7, 8}
	// for _, v := range arr {
	// 	mun = mun + v
	// }
	// for i, v := range arr {
	// 	for j, n := range arr {
	// 		if v+n == 8 {
	// 			fmt.Println(i, j)
	// 		}
	// 	}
	// }
	// // 切片(slice)
	// // 切片的定义
	// var s1 []int    //定义了一个存放int类型的切片,在没有赋值的时候是空
	// var s2 []string //定义了一个存放string类型的切片
	// s1 = []int{1, 2, 3}
	// s2 = []string{"w", "z", "c"}
	// fmt.Println(s1, s2)

	// var s3 = make([]int, 2)
	// fmt.Println(s3)
}
