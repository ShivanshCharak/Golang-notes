package main

import (
	"fmt"
	"time"
)

func someFunc(number int) {
	fmt.Println(number)
}
func main() {
	go someFunc(1)
	time.Sleep(time.Second * 2)
	fmt.Println("hi")
}
