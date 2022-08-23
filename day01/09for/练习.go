package main

import "fmt"

func forMain() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i >= j {
				fmt.Printf("%d*%d=%d  ", i, j, i*j)
			}
		}
		fmt.Printf("\n")
	}
}
