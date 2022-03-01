package main

import "fmt"

type Movement interface {
	move()
}

type Animal interface {
	speak()
}

type NextAnimal interface {
	Movement
	Animal
}

type data struct {
	index int
}

type Dog struct{}

func (d Dog) speak() {
	fmt.Println("woaww")
}

func (d Dog) move() {
	fmt.Println("chay bon chan")
}

func goout(i interface{}) {
	fmt.Println(i)
}

func main() {
	var animal Animal
	animal = Dog{}

	animal.speak()

	dog := Dog{}
	var m Movement = dog
	m.move()

	var a Animal = dog
	a.speak()

	var nextAnimal = dog
	nextAnimal.move()
	nextAnimal.speak()

	goout(10)
	goout(10.12)

	d := data{123}
	goout(d)

}
