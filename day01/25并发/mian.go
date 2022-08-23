package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello")
}

func main() {
	fmt.Println("00")
	go hello()
	fmt.Println("01")
	time.Sleep(1 * time.Second)
}
