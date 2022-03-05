package main

import (
	"fmt"
	"time"
)

func forLoop(str string) {
	for i := 0; i < 1000; i++ {
		fmt.Println(str, ":", i)
	}
}

func main() {
	//Call loop in a goroutine
	go forLoop("direct")

	//Call another loop in another goroutine
	go forLoop("goroutine")

	//Call another another function in another goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//Sleep to wait for goroutine to end
	time.Sleep(time.Second)

	//Print the last message
	fmt.Println("done")
}
