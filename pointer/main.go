package main

import "fmt"

func main() {
	a := 100
	var pointer *int
	pointer = &a
	fmt.Println(pointer)

	fmt.Printf("%T", pointer)
	fmt.Println()
	fmt.Printf("%p", pointer)
	fmt.Println()
	fmt.Printf("%p", &pointer)
	fmt.Println()

	b := 123
	p2 := new(int)

	p2 = &b

	fmt.Println(p2)
	fmt.Printf("%T", pointer)
	fmt.Println()

	var p3 *int
	fmt.Println(p3)

	p5 := new(int)
	fmt.Println(p5)

	var p6 *int
	c := 100
	p6 = &c

	fmt.Println(c)
	*p6 = 999

	fmt.Println("p6", p6)
	fmt.Println(c)

	array := [3]int{1, 2, 3}
	var p7 *[3]int
	p7 = &array

	fmt.Println("p7", *p7)
	fmt.Println(p7)
	fmt.Println(&p7)
	fmt.Printf("%T", p7)
	fmt.Println()

	d := 30
	var p8 *int = &d
	applyPointer(p8)

	fmt.Println()
	fmt.Println(d)

}

func applyPointer(pointer *int) {
	// println(pointer, "pointer")
	*pointer = 777
}
