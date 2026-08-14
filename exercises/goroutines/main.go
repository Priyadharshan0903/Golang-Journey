package main

import (
	"fmt"
	"time"
)

func sendConfirmation(order string) string {
	time.Sleep(500 * time.Millisecond)
	return "Confirmed" + order
}

func main() {
	reply := make(chan string)

	fmt.Println("started")

	go func() {
		reply <- sendConfirmation(" No. 2")
	}()

	// Go has no `typeof`. %T is a formatting verb, so it needs Printf, not Println.
	// (fmt.Println(reply) would print the channel's address instead.)
	fmt.Printf("reply is %T\n", reply) // chan string

	// The channel type and its element type are different things.
	fmt.Printf("<-reply is %T\n", <-reply) // string

	fmt.Println("ended")

	channelDemo()
	waitGroupDemo()
	selectDemo()
}
