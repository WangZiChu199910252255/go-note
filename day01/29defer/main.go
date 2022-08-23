package main

import (
	"fmt"
	"os"
)

func test() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func test2() {
	f, err := os.Open("./a.txt")
	if err != nil {
		defer f.Close()
	} else {
		fmt.Println(err)
	}
}

func main() {

}
