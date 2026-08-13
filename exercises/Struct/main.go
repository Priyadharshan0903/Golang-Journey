package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func main() {
	r := Rectangle{Width: 5, Height: 5}
	fmt.Println(r.Area())
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// only we can read the values ,
// we can't edit it since the copy of the variable is passed
//
// In order to edit the params or arguements passed ,
// we must send the Pointer* of it

// Pointer Receiver
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}
