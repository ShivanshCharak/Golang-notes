# Concurrency

## GO routines

```package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}
// fork join model
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
```
- Go routines are just like async function in js 
- Go keywork tells the function to fork it and and we have to use the time.sleep method which tell the function to join it back with the main function

![Fork join model](./fork-join-model.png)
