package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*1000)

	req, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	fmt.Println("Respose received, status code", res.StatusCode)
}
