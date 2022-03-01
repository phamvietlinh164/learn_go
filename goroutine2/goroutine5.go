package main

import "fmt"

func main() {
	c := make(chan string)
	c <- "hello"

	msg := <-c
	fmt.Println(msg)

	// chương trình trên báo lỗi vì send chữ hello vào channel c sẽ bị
	// đứng ngay tại đó cho tới khi 1 go routine khác lấy nó ra
	// chương trình không chạy tới được dòng tiếp theo
}
