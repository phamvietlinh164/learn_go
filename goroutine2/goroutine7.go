package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2s"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	// chương trình này chạy print c1, sau đó phải đợi 2s mới print c2
	// mặc dù c1 chỉ cần 500ms đã đưa vào channel
	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}
