package main

import "fmt"

func test() {
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Printf("m1: %v\n", m1)
}

func test1() {
	var m1 = map[string]string{"name": "wzc", "age": "20"}
	m2 := make(map[string]string)
	m2["name"] = "wzc"
	m2["age"] = "18"
	fmt.Printf("m1: %v\n m2: %v\n", m1, m2)
}

func test2() {
	var m1 = map[string]string{"name": "wzc", "age": "20"}
	var key = "name"
	value := m1["name"]
	fmt.Printf(" key: %v\n value: %v", key, value)
}

func test3() {
	var m1 = map[string]string{"name": "wzc1", "age": "20"}
	var k1 = "name"
	var k2 = "age1"
	v, ok := m1[k1]
	v2, ok2 := m1[k2]
	fmt.Printf("k1: %v v1: %v ok: %v\n", k1, v, ok)
	fmt.Printf("k2: %v v2: %v ok2: %v\n", k2, v2, ok2)
}

func test4() {
	m1 := map[string]string{"a": "b", "c": "d", "e": "f"}
	for k, v := range m1 {
		fmt.Printf("k: %v v: %v\n", k, v)
	}
	fmt.Printf("m1: %T\n", m1)
}
func main() {
	// test()
	// test1()
	// test2()
	// test3()
	test4()
}
