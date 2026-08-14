package main

import (
	"fmt"
	"time"
)

func selectDemo() {
	fmt.Println("\n--- select ---")

	fast := make(chan string)
	slow := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		fast <- "fast finished"
	}()
	go func() {
		time.Sleep(150 * time.Millisecond)
		slow <- "slow finished"
	}()

	// select blocks until ONE case is ready. Think Promise.race — except a
	// Promise.race is one-shot, while this sits in a loop and keeps racing.
	// If several cases are ready at once, Go picks one at random.
	for range 2 {
		select {
		case msg := <-fast:
			fmt.Println("got:", msg)
		case msg := <-slow:
			fmt.Println("got:", msg)
		}
	}

	// TIMEOUT: time.After returns a channel that fires after the duration,
	// so it's just another case. This is Go's answer to Promise.race([p, sleep]).
	quiet := make(chan string)
	select {
	case msg := <-quiet:
		fmt.Println("got:", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out waiting on quiet")
	}

	// NON-BLOCKING: a default case runs immediately when nothing else is ready,
	// turning select into a poll instead of a wait.
	select {
	case msg := <-quiet:
		fmt.Println("got:", msg)
	default:
		fmt.Println("nothing ready, moving on")
	}
}
