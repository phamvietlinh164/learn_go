package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	cxt, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	mySleepAndTalk(cxt, 5*time.Second, "hello")

}

func mySleepAndTalk(cxt context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-cxt.Done():
		fmt.Println(cxt.Err())
	}
}
