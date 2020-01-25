package main

import (
	"context"
)

func DoNotLeak(ctx context.Context, send chan int) chan int {
	recv := make(chan int)
	go func() {
		for {
			select {
			case num := <-recv:
				send <- num
			case <-ctx.Done():
				return
			}
		}
	}()
	return recv
}

func DoLeak(send chan int) chan int {
	recv := make(chan int)
	go func() {
		for {
			select {
			case num := <-recv:
				send <- num
			}
		}
	}()
	return recv
}
