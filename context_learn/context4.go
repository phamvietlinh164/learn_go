package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	time.AfterFunc(time.Second*2, func() {
		cancel()
	})

	select {
	case <-ctx.Done():
		fmt.Println("Done")
	}
}
