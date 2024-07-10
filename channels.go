package main

import "fmt"

func main() {
	myChannel := make(chan int)
	go func() {
		add := 3 + 4*7 + 9
		myChannel <- add
	}()
	//  invoking now cuz its a annonymous function
	msg := <-myChannel

	fmt.Println(msg)
	fmt.Println("hey")
}
