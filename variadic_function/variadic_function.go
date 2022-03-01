package main

import "fmt"

func addItem(item int, list ...int) {
	list = append(list, item)
	fmt.Println(list)
}

func main() {
	addItem(1, 100, 200, 300)
}
