package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	count("fish")

	// Nếu tạo 2 go routine cùng chạy thì không nhận được kết quả vì
	// chương trình sẽ chạy qua 2 dòng code rồi kết thúc
	// go count("sheep")
	// go count("fish")

}

func count(thing string) {
	for i := 1; true; i = i + 1 {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
