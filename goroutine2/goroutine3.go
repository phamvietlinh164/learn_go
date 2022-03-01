package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)
	// chương trình sẽ báo lỗi vì khi hàm count kết thúc, msg vẫn chờ lấy giá trị của channel
	// Lỗi này xảy ra ở run time, không phải ở build time
	for {
		msg := <-c

		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i = i + 1 {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
}
