package main

import (
	"fmt"
	"strconv"
	"time"

)
func goroutine(){
	for i:=1;i<=10;i++{
		fmt.Println("hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main(){
	// 开启一个协程
	go goroutine()

	for i:=1;i<=10;i++{
		fmt.Println("hello goroutine" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
