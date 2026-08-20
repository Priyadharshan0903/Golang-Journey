package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// context is typically for timeouts
// cancelling go routines
// and passing metadata across the application
func main() {
	fmt.Println("Using Contexts")
	// ctx := context.Background()
	// exampleTimeout()

	// exampleWithValues()

	http.HandleFunc("/hello", Handler)
	http.ListenAndServe(":8080", nil)

}

func exampleTimeout() {
	ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Called the API!")
	case <-ctxWithTimeOut.Done():
		fmt.Println("Oh My Timeout has expired!", ctxWithTimeOut.Err())
		// we will do something later here

	}
}

func exampleWithValues() {
	type key int
	const UserKey key = 0
	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx, UserKey, "123")

	if userId, ok := ctxWithValue.Value(UserKey).(string); ok {
		fmt.Println("User Id : ", userId)
	} else {
		fmt.Println("This is the protected route - no userId found")
	}
}
