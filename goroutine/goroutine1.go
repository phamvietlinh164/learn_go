package main

import (
	"fmt"
	"time"
)

func g1() {
	for i := 0; ; i = i + 1 {

		fmt.Println("g1", i)
		time.Sleep(time.Second)
	}
}

func g2() {
	for i := 0; ; i = i + 1 {
		fmt.Println("g2", i)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	go g1()
	go g2()

	var input string
	fmt.Scanln(&input)
}
