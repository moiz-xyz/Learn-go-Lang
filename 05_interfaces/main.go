package main

import (
	"fmt"
	"math"
)

// 1. Define the interface (The Contract)
type Shape interface {
	Area() float64
}

// 2. Create a Circle struct
type Circle struct {
	Radius float64
}

// 3. Circle implements the Shape interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 4. Create a Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// 5. Rectangle implements the Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 6. A function that accepts ANY shape
func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	// We can pass both structs to PrintArea because both implement Shape
	PrintArea(c) 
	PrintArea(r) 
}


package main

import (
	"fmt"
	"math"
)

// 1. Define the interface (The Contract)
type Shape interface {
	Area() float64
}

// 2. Create a Circle struct
type Circle struct {
	Radius float64
}

// 3. Circle implements the Shape interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 4. Create a Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// 5. Rectangle implements the Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 6. A function that accepts ANY shape
func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	// We can pass both structs to PrintArea because both implement Shape
	PrintArea(c) 
	PrintArea(r) 
}


// Multiple Types Can Use the Same Interface
type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func main() {
	var s Speaker

	s = Dog{}
	s.Speak()

	s = Cat{}
	s.Speak()
}