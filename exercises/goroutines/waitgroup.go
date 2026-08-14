package main

import (
	"fmt"
	"sync"
	"time"
)

func waitGroupDemo() {
	fmt.Println("\n--- waitgroup ---")

	var wg sync.WaitGroup

	// Each goroutine writes to its own index, so no mutex is needed —
	// distinct slice slots are independent memory.
	results := make([]string, 3)

	for i := range results {
		// wg.Go (Go 1.25+) does Add(1), starts the goroutine, and calls Done
		// for you. `i` is per-iteration since Go 1.22, so no shadow copy needed.
		wg.Go(func() {
			time.Sleep(time.Duration(i*50) * time.Millisecond)
			results[i] = fmt.Sprintf("worker %d done", i)
		})
	}

	// Blocks until the counter hits zero. Roughly Promise.all, except it
	// waits for completion rather than collecting return values —
	// goroutines can't return anything, which is why we wrote into `results`.
	wg.Wait()

	for _, r := range results {
		fmt.Println(r)
	}

	// The older form you'll see everywhere, equivalent to one wg.Go call:
	//
	//     wg.Add(1)
	//     go func() {
	//         defer wg.Done() // defer so it runs even if the body panics
	//         work()
	//     }()
	//
	// A WaitGroup must not be copied after first use — pass *sync.WaitGroup,
	// never sync.WaitGroup, or Done() will decrement the wrong counter.
}
