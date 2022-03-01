package main

import (
	"fmt"
	"time"
)

func createPizza1(pizza int) {
	time.Sleep(time.Second)
	fmt.Printf("Create pizza %d \n", pizza)
}

func timeTrack1(start time.Time, functionName string) {
	elapsed := time.Since(start)
	fmt.Println(functionName, "took", elapsed)
}

func main() {
	defer timeTrack1(time.Now(), "Build pizzas")

	for i := 0; i < 3; i = i + 1 {
		createPizza1(i)
	}
}
