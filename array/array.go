package main

import "fmt"

func main() {
	var myArray [4]int

	fmt.Println(myArray)

	// gan gia tri cho array
	myArray[0] = 123

	fmt.Println(myArray)

	myArray[3] = 456

	fmt.Println(myArray)

	// khai bao 1 array co khoi tao gia tri
	arrays := [3]int{1, 2, 3}
	fmt.Println(arrays)

	arrays2 := [3]int{100}
	fmt.Println(arrays2)

	// size mang
	fmt.Println(len(arrays2))

	// khai bao mang khong can size
	array3 := [...]string{"Vinfast", "Honda", "Huyndai", "Tesla"}
	fmt.Println(array3)
	fmt.Println(len(array3))

	// array la value type khong phai reference type
	contries := [...]string{"VN", "US", "FRANCE"}
	copyContries := contries

	copyContries[0] = "CANADA"

	fmt.Println(contries)
	fmt.Println(copyContries)

	// loop array
	// cach 1
	for i := 0; i < len(contries); i = i + 1 {
		fmt.Println(contries[i])
	}

	// cach 2
	for index, value := range contries {
		fmt.Printf("i=%d value=%s", index, value)
		fmt.Println()
	}

	// mang hai chieu
	arr := [4][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}
	fmt.Println(arr)
	for i := 0; i < 4; i = i + 1 {
		for j := 0; j < 2; j = j + 1 {
			fmt.Print(arr[i][j], " ")
		}
		println()
	}
}
