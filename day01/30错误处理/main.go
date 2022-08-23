package main

import (
	"errors"
	"fmt"
)

func main() {
	panic("panic error")
	err := errors.New("error")
	fmt.Println(err)
}
