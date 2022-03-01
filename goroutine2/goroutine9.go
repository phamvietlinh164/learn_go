package main

import "fmt"

func main() {
	jobs := make(chan int, 10)
	result := make(chan int, 10)

	go worker(jobs, result)
	go worker(jobs, result)

	for i := 0; i < 10; i = i + 1 {
		jobs <- i
	}
	close((jobs))

	for j := 0; j < 10; j = j + 1 {
		fmt.Println(<-result)
	}

}

// định nghĩa job channel chỉ có thể receive, result channel chỉ có thể send
func worker(job <-chan int, result chan<- int) {
	for n := range job {
		result <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
