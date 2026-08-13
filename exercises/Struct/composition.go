package main

import "fmt"

// Go has no inheritance and no `extends`. Instead you EMBED one struct
// inside another, and the outer type gets the inner one's fields and methods.

type Shape struct{ Name string }

func (s Shape) Describe() {
	fmt.Println("This is", s.Name)
}

// Box embeds Shape. Embedding = a type written with NO field name and NO comma.
// Writing `Shape,` would declare a field *named* Shape of the following type
// instead, which is a plain field, not embedding.
type Box struct {
	Shape
	Width, Height float64
}

func (b Box) Area() float64 {
	return b.Width * b.Height
}

func compositionDemo() {
	// The embedded field's name is its type name, so you set it as Shape:
	b := Box{
		Shape:  Shape{Name: "Box"},
		Width:  4,
		Height: 3,
	}

	// Describe() is PROMOTED from Shape — Box never declares it
	b.Describe()

	// Fields are promoted too: b.Name is shorthand for b.Shape.Name
	fmt.Println("promoted:", b.Name, "| explicit:", b.Shape.Name)

	// Box's own method sits alongside the promoted ones
	fmt.Println("area:", b.Area())

	// Outer wins: if Box declared its own Describe(), it would shadow Shape's,
	// and the embedded one would still be reachable at the explicit path.
	b.Shape.Describe()
}
