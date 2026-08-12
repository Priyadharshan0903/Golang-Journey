package main

import "fmt"

func main() {
	condition := true

	if condition {
		fmt.Println("Pass Case")
	}

	val := false

	if val {
		fmt.Println("True Case")
	} else {
		fmt.Println("Else Case")
	}
}
