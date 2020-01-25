package main

import (
	"context"
	"runtime"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestDoNoLeak(t *testing.T) {
	startNum := runtime.NumGoroutine()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	out := make(chan int)
	in := DoNotLeak(ctx, out)
	in <- 10

	res := <-out

	time.Sleep(1 * time.Second)
	endNum := runtime.NumGoroutine()
	t.Logf("goroutine num: %d -> %d\n", startNum, endNum)

	if res != 10 {
		t.Errorf("want 10, got %d", res)
	}
}

func TestDoLeak(t *testing.T) {
	startNum := runtime.NumGoroutine()

	out := make(chan int)
	in := DoLeak(out)
	in <- 10

	res := <-out

	time.Sleep(1 * time.Second)
	endNum := runtime.NumGoroutine()
	t.Logf("goroutine num: %d -> %d\n", startNum, endNum)

	if res != 10 {
		t.Errorf("want 10, got %d", res)
	}
}
