package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello da!")
}

func main() {

	fmt.Println("Start!")
	go sayHello()
	fmt.Println("Midd")
	// if comment out the below line, you won't see the output from the sayHello goroutine
	time.Sleep(100 * time.Millisecond) // hacky way , should use either channels or waitGroups

	fmt.Println("Now Using WaitGroups")
	WithWgGroups()

	fmt.Println("----------------------------------")
	fmt.Println("Now using Channels")
	WithChannels()
}
