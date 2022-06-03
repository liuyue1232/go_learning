package csp

import (
	"fmt"
	"time"
)

func Hello(n int) int {
	// c := make(chan int)
	// c <- 12
	time.Sleep(2 * time.Second)
	// ch := <-c
	fmt.Println(5)

	return n + 1
}
