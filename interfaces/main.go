package main

import (
	"fmt"
	"math"
	"strings"
)

type Square struct {
	side float64
}

// By creating this function, we implement the
// Shape interface, since it declared the area
// function. Don't forget, a square is a shape.
func (square Square) area() float64 {
	return square.side * square.side
}

type Circle struct {
	radius float64
}

// By creating this function, we implement the
// Shape interface, since it declared the area function.
// Don't forget, a circle is a shape.
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

type Shape interface {
	// Every struct who has this method "area" with the exact
	// same interface, automatically implements this shape interface.
	area() float64
}

// With this function we can print info
// of a shape.
func info(shape Shape) {
	fmt.Println(strings.Repeat("*", 40))
	fmt.Printf("The type of this Shape is: '%T'\n", shape)
	fmt.Println("The calculated area is:", shape.area())
}

func main() {
	// Since a square and a circle is a shape,
	// we can use the info function on both
	// types. This is also called polymorphism
	square := Square{10}
	circle := Circle{5}
	info(square)
	info(circle)
}
