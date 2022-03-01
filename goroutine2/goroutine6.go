package main

import "fmt"

func main() {
	c := make(chan string, 2)
	// có thể định capacity của channel
	// khi capicity của channel là 2 thì có thể đưa 2 giá trị vào channel
	// khi định capacity của channel là 1 thì có thể đưa 1 giá trị
	// vào channel mà không bị block
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
