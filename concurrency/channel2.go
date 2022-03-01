package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 22

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Khi close channel thi gia tri lay ra tu channel bang 0
	ch2 := make(chan int, 2)

	ch2 <- 111
	ch2 <- 222
	close(ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

}
