package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type Email struct{ addr string }
type SMS struct{ number string }

func (e Email) Send(message string) error {
	fmt.Println("This is the Email Struct's reference")
	return nil
}

func (e SMS) Send(message string) error {
	fmt.Println("This is the SMS Struct's reference")
	return nil
}

func notify(n Notifier, message string) error {
	return n.Send(message)
}

// Always  we can't put executables outside i.e at package level
// we must include them inside a function
// eg: The notify (...) can be called outside the function main in JS
// but in go everything needs to be inside function since it is package based

func main() {
	notify(Email{addr: "priyadharshansenthil"}, "Mosi Mosi")
	notify(SMS{}, "Gambare")
}
