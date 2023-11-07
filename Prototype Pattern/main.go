package main

import (
	"fmt"
	"prototype-pattern/shapes"
)

func main() {
	originalCircle := &shapes.Circle{Radius: 1}
	fmt.Println("First Circle Area: ", originalCircle.GetArea())
	fmt.Println("First Circle Perimeter: ", originalCircle.GetPerimeter())

	newCircle := originalCircle.Clone()
	newCircle.SetMeasurements(10)
	fmt.Println("Second Circle Area: ", newCircle.GetArea())
	fmt.Println("Second Circle Perimeter: ", newCircle.GetPerimeter())

	fmt.Println()

	originalRectangle := &shapes.Rectangle{Length: 1, Breadth: 1}
	fmt.Println("First Rectangle Area: ", originalRectangle.GetArea())
	fmt.Println("First Rectangle Perimeter: ", originalRectangle.GetPerimeter())

	newRectangle := originalRectangle.Clone()
	newRectangle.SetMeasurements(10, 10)
	fmt.Println("Second Rectangle Area: ", newRectangle.GetArea())
	fmt.Println("Second Rectangle Perimeter: ", newRectangle.GetPerimeter())
}
