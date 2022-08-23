package main

import (
	"fmt"
)

type Person struct {
	id    int
	name  string
	age   int
	email string
}

type User struct {
	name string
}

func (u *User) setName(name string) {
	u.name = name
}

func test() {
	var p_person *Person
	tom := Person{
		id:    101,
		name:  "Tom",
		age:   20,
		email: "tom@gmail.com",
	}
	p_person = &tom
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %p\n", p_person)
	fmt.Printf("p_person: %v\n", *p_person)
}

func test1() {
	// tom := Person{
	// 	id:    101,
	// 	name:  "Tom",
	// 	age:   20,
	// 	email: "tom@gmail.com",
	// }
	tom := new(Person)
	tom.id = 101
	tom.name = "Tom"
	fmt.Printf("tom: %v\n", *tom)
}

func showPerson(person Person) Person {
	person.id = 101
	person.name = "Tom"
	fmt.Printf("person: %v\n", person)
	return person
}

func showPerson2(person *Person) {
	person.id = 101
	person.name = "Tom"
	fmt.Printf("person: %v\n", *person)
}

func test2() {
	type Dog struct {
		name  string
		age   int
		color string
	}
	type Person struct {
		dog  Dog
		name string
		age  int
	}
	dog := Dog{name: "花花", age: 2, color: "red"}
	per := Person{dog: dog, name: "tom", age: 20}
	fmt.Printf("per: %v\n", per)
}
func (per Person) eat() {
	fmt.Printf("%v,eat", per.name)
}
func (per Person) sleep() {
	fmt.Printf("%v,sleep", per.name)
}
func test3() {
	per := Person{
		id:    101,
		name:  "Tom",
		age:   20,
		email: "tom@gmail.com",
	}
	per.eat()
	per.sleep()
}

func main() {
	user := User{
		name: "Tom",
	}
	user.setName("tom")
	fmt.Printf("user: %v\n", user)
	// per := person.People{}
	// test3()
	// test2()
	// test1()
	// var tom struct {
	// 	id   int
	// 	name string
	// 	age  int
	// }
	// var tom Person
	// tom.id = 101
	// tom.name = "tom"
	// tom.age = 20
	// tom.email = "tom@gmail.com"
	// tom := Person{
	// 	id:    101,
	// 	name:  "Tom",
	// 	age:   20,
	// 	email: "tom@gmail.com",
	// }
	// tom := Person{
	// 	1,
	// 	"name",
	// 	20,
	// 	"tom@gmail.com",
	// }
	// // person := &tom
	// tom = showPerson(tom)
	// fmt.Printf("tom: %v\n", tom)

}
