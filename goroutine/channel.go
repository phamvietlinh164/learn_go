package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	// time.Sleep(time.Second * 3)
	for i := 0; i < 10; i = i + 1 {
		fmt.Println("truoc khi dua vao channel")
		c <- "ping"
		fmt.Println("sau khi dua vao channel")
		// fmt.Println("abc")
		time.Sleep(time.Second * 2)
		// fmt.Println("sau khi sleep")
	}
}

func printer(c chan string) {
	time.Sleep(time.Second * 10)
	for {

		fmt.Println("truoc khi lay tu channel")
		msg := <-c
		fmt.Println(msg)
		fmt.Println("sau khi lay tu channel")
		// time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)

}
