package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("API Response!")
	case <-ctx.Done():
		fmt.Println("Oh The Context has expired")
		http.Error(w, "Request Context Timeout", http.StatusRequestTimeout)
		return
	}

}
