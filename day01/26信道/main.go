package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(done chan bool) {
	fmt.Println("wzc")
	done <- true
}

func test1(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	fmt.Println("aaa")
	close(c)
}

func test2(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	time.Sleep(1 * time.Second)
	wg.Done()
}

func main() {
	// done := make(chan bool)
	// go hello(done)
	// <-done
	// a := make(chan int)
	// go test1(a)
	// for {
	// 	v, ok := <-a
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go test2(i, &wg)
	}
	wg.Wait()
}
