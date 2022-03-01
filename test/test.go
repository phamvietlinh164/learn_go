package main

import (
	"fmt"
	"time"
)

func main() {

	// Declaring Time in UTC
	Time := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)
	// Calling AddDate method with all
	// its parameters
	t1 := Time.AddDate(1, 2, 5)
	t2 := Time.AddDate(5, -2, 9)
	t3 := Time.AddDate(0, 3, -3)
	t4 := Time.AddDate(1, 0, 0)

	// Prints output
	fmt.Printf("%v\n", Time)
	fmt.Printf("%v\n", t1)
	fmt.Printf("%v\n", t2)
	fmt.Printf("%v\n", t3)
	fmt.Printf("%v\n", t4)
}
