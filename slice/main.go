package main

import (
	"fmt"
)

func main() {
	var slice []int
	fmt.Println(slice)

	// khai bao va khoi tao slice
	var slice1 = []int{1, 2, 3, 4}
	fmt.Println(slice1)

	// tao mot slice tu 1 array
	var array = [4]int{1, 2, 3, 4}
	slice2 := array[1:3]
	fmt.Println(slice2)

	slice3 := array[:]
	fmt.Println(slice3)

	slice5 := array[2:]
	fmt.Println(slice5)

	// tao mot slice tu mot slice khac
	var slice6 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice7 := slice6
	fmt.Println(slice7)

	slice8 := slice6[1:]
	fmt.Println(slice8)

	// slice la mot reference cua array
	var array1 = [5]int{1, 2, 3, 4, 5}
	slice9 := array1[:]
	slice9[0] = 777

	fmt.Println(slice9)
	fmt.Println(array1)

	// length and capacity
	contries := [...]string{"VN", "USA", "CANADA", "FRANCE", "CHINA"}
	slice10 := contries[2:3]
	fmt.Println(slice10)

	fmt.Println(len(slice10))
	fmt.Println(cap(slice10))

	// make copy append
	slice11 := make([]int, 2, 5)
	fmt.Println(slice11)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))

	slice12 := make([]int, 2)
	fmt.Println(slice12)
	fmt.Println(len(slice12))
	fmt.Println(cap(slice12))

	var slice15 []int
	slice15 = append(slice15, 100)
	fmt.Println(slice15)

	src := []string{"A", "B", "C", "D"}
	dest := make([]string, 2)

	number := copy(dest, src)
	fmt.Println(number)
	fmt.Println(dest)

	s := []int{2, 3, 5, 7, 11, 13}
	// printSlice(s)

	// Slice the slice to give it zero length.
	s = s[1:3]
	fmt.Printf("%v \n", s)

	// Extend its length.
	s = s[:4]
	fmt.Printf("%v \n", s)

	// Drop its first two values.
	// s = s[2:]
	// fmt.Printf("%v \n", s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
