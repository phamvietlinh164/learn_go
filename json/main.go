package main

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Id    string `json:"name"`
	Title string `json:"age"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

func main() {
	data := []byte(`{"name": "Linh", "age": "28"}`)

	var app App
	// fmt.Println(a[1])``
	err := json.Unmarshal(data, &app)

	if err == nil {
		fmt.Println(app)
	} else {
		fmt.Println(err)
	}

	social := Social{Facebook: "https://facebook.com", Twitter: "https://twitter.com"}
	user := User{Name: "LanKa", Type: "Author", Age: 25, Social: social}

	byteArray, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

}
