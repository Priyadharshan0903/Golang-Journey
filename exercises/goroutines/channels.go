package main

import "fmt"

func channelDemo() {
	fmt.Println("\n--- channels ---")

	// UNBUFFERED: a rendezvous, not a queue. The send blocks until some other
	// goroutine is ready to receive, and the receive blocks until someone sends.
	// Unlike a JS Promise (resolves once, caches, any number of .then readers),
	// each value here is handed to exactly one receiver.
	unbuffered := make(chan string)
	go func() { unbuffered <- "handed off" }()
	fmt.Println("unbuffered:", <-unbuffered)

	// BUFFERED: sends don't block until the buffer is full, so these two sends
	// need no goroutine at all. len/cap work like they do on slices.
	buffered := make(chan string, 2)
	buffered <- "first"
	buffered <- "second"
	fmt.Printf("buffered: len=%d cap=%d\n", len(buffered), cap(buffered))

	// close() says "no more values coming". Ranging a channel drains it until
	// closed — same `range` keyword as the slices exercise, but it yields
	// values only, never an index.
	close(buffered)
	for v := range buffered {
		fmt.Println("range chan:", v)
	}

	// comma-ok: ok is false once the channel is closed AND drained.
	v, ok := <-buffered
	fmt.Printf("drained: v=%q ok=%v\n", v, ok)

	// DEADLOCK — do not uncomment. An unbuffered send with no other goroutine
	// running blocks forever, and the runtime notices every goroutine is stuck:
	//
	//     stuck := make(chan string)
	//     stuck <- "nobody is listening"
	//
	//     fatal error: all goroutines are asleep - deadlock!
	//
	// This is a runtime fatal error, not a panic — recover() cannot catch it.
}
