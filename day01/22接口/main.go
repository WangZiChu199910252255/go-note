package main

import "fmt"

type USB interface {
	read()
	write()
}

type Computer struct {
	name string
}

type Mobile struct {
	model string
}

func (m Mobile) read() {
	fmt.Printf("m.name: %v", m.model)
}

func (m Mobile) write() {
	fmt.Printf("m.model: %v\n", m.model)
}

func (c Computer) read() {
	fmt.Printf("c.name: %v", c.name)
	fmt.Printf("read...\n")
}

func (c Computer) write() {
	fmt.Printf("c.name: %v", c.name)
	fmt.Printf("write...\n")
}

type Pet interface {
	eat()
}

type Dog struct {
}

type Cat struct {
}

func (dog Dog) eat() {
	fmt.Printf("dog eat \n")
}

func (cat Cat) eat() {
	fmt.Printf("cat eat")
}

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct {
}

func (fish Fish) fly() {
	fmt.Printf("fish fly")
}

func (fish Fish) swim() {
	fmt.Printf("fish swim")
}
func main() {
	// c := Computer{name: "联想"}
	// c.read()
	// c.write()

	// m := Mobile{
	// 	model: "5G",
	// }
	// m.read()
	// m.write()
	// var pet Pet
	// pet = Dog{}
	// pet.eat()
	// pet = Cat{}
	// pet.eat()
	var ff FlyFish = Fish{}
	ff.fly()
	ff.swim()
}
