package main

import (
	"encoding/json"
	"fmt"
)

type Secret struct {
	Message string
}

func main() {
	sec := Secret{"I love you!"}

	byte1, err := json.Marshal(sec)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T", byte1)
	raw := json.RawMessage(byte1)
	fmt.Printf("Raw Message : %s\n", raw)
}
