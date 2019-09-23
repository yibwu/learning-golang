package main

import "fmt"

type Animal interface {
	sayHello()
}

type Dog struct {
}

func (d Dog) sayHello() {
	fmt.Println("wang~")
}

type Cat struct {
}

func (c Cat) sayHello() {
	fmt.Println("miao~")
}

func SayHello(animal Animal) {
	animal.sayHello()
}

func main() {
	//var animal Animal
	//
	//animal = Dog{}
	//animal.sayHello()
	//
	//animal = Cat{}
	//animal.sayHello()

	dog := Dog{}
	cat := Cat{}
	SayHello(dog)
	SayHello(cat)
}
