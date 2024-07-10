package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}
func main() {
	go someFunc("1")
	go someFunc("5")
	go someFunc("4")
	go someFunc("3") // async fork the fnction from the main
	// here it is not rejoinsing with main but if we use time.sleep it will rejoin with the main function
	// go is juts like async in golang these ae goroutines
	time.Sleep(time.Second * 2)

	fmt.Println("hey")
}
