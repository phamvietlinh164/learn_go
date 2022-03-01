package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	for {
		msg, open := <-c

		if !open {
			break
		}
		fmt.Println(msg)

	}

	// hoặc có thể viết như sau để không cần check open
	// for msg := range c {

	// 	fmt.Println(msg)

	// }
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i = i + 1 {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
