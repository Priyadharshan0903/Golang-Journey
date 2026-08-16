package main

import (
	"fmt"
	"sync"
	"time"
)

// Same package as main.go, so Order is already in scope — no import, no redeclaring.
//
// Note there's no *sync.WaitGroup parameter here: wg.Go owns the Add/Done pair
// itself, so the worker stays a plain function that knows nothing about waiting.
func serveOrder(e Order) {
	fmt.Println("Executing the Order....", e.TableNumber)
	time.Sleep(e.PrepTime)
	fmt.Printf("Order ready for %d \n", e.TableNumber)

	fmt.Println("----------------------------------------")
}

func runWithWaitGo(orders []Order) {
	fmt.Println("--- WaitGroup.Go version ---")

	var wg sync.WaitGroup

	for _, order := range orders {
		// wg.Go does the wg.Add(1) and the deferred wg.Done() for you.
		// `order` is captured straight from the loop: since Go 1.22 each
		// iteration gets its own variable, so no `order := order` shadow.
		wg.Go(func() { serveOrder(order) })
	}

	wg.Wait()
}
