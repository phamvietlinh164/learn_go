package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var myMap = make(map[string]int)
	fmt.Println(myMap)

	if myMap == nil {
		fmt.Println("myMap = nil")
	} else {
		fmt.Println("myMap != nil")
	}

	var myMap1 map[string]int
	fmt.Println(myMap1)

	if myMap1 == nil {
		fmt.Println("myMap1 = nil")
	}

	// khai bao voi gia tri khoi tao
	myMap2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}

	fmt.Println(myMap2)

	// them 1 phan tu vao map
	myMap2["key4"] = 4
	myMap2["key5"] = 5

	fmt.Println(myMap2)

	// delete 1 phan tu trong map
	delete(myMap2, "key1")
	fmt.Println(myMap2)

	// map la 1 reference type
	myMap3 := myMap2
	fmt.Println(myMap3)

	myMap3["key5"] = 1000
	fmt.Println(myMap2)

	// truy cap mot phan tu trong map
	value, found := myMap2["key1000"]
	fmt.Println(value)

	if found {
		fmt.Println(value)
	} else {
		fmt.Println("Khong tim thay gia tri")
	}

	var myMap5 = make(map[string]interface{})
	var myMap6 = make(map[string]string)

	myMap6["a"] = "aa"
	myMap6["b"] = "bb"
	var sli = []string{"s", "k"}
	type Aa struct {
		a string
		b int
	}

	a := Aa{"aaa", 1}

	myMap5["abc"] = "abcd"
	myMap5["a"] = 1
	myMap5["map"] = myMap6
	myMap5["sli"] = sli
	myMap5["str"] = a
	jsonString, err := json.Marshal(myMap5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myMap5)
	fmt.Println(string(jsonString))

	a1 := make(map[int]string)

	a1[1] = "John"

	j, err := json.Marshal(a1)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(j))
	}

}
