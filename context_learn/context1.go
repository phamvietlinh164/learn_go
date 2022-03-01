package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	fmt.Println("abc")
}
