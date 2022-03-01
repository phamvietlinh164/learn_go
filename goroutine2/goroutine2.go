package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// wg add chỉ định bao nhiêu hàm Done cần phải call để kết thúc wait, sau đó mới chạy tiếp chương trình
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("end")
}

func count(thing string) {
	for i := 1; i <= 5; i = i + 1 {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
