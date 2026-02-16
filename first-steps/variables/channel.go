package main

import (
	"fmt"
	"time"
)

func DemoChannel() {
	fmt.Println("== Channel demo ==")
	ch := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- "ping"
	}()
	msg := <-ch
	fmt.Println("received", msg)
}
