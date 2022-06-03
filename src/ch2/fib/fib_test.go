package fib

import (
	"fmt"
	"testing"
)

func TestFibFirt(t *testing.T) {
	var a int = 1
	var b int = 1
	fmt.Println(" ", a)

	for i := 0; i < 5; i++ {
		fmt.Println(" ", b)
		temp := a
		a = b
		b = temp + a
	}
	fmt.Println()
}
