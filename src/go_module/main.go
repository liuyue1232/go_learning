package main

import "github.com/sirupsen/logrus"

func main() {
	var {
        a = 13
        b = int32(17)
        f = float32(3.14)
    }
    logrus.Println("hello, gopath mode")
}
