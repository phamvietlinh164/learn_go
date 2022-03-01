package main

import "fmt"

func main() {
	queue := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i = i + 1 {
			queue <- i
		}
		close(queue)
	}()

	for value := range queue {
		fmt.Println(value)
	}

}
