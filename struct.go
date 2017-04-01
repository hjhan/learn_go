package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

/**
Function
*/
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

/**
Methods  by added a "receiver"
*/
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	/**
	Structs
	*/
	c := Circle{x: 0, y: 0, r: 5}

	fmt.Println(c.x, c.y, c.r)

	fmt.Println(circleArea(c))

	fmt.Println(c.area())

	fmt.Printf("hello %s\n", "world!")
	fmt.Printf("Sqrt(9)=%f", math.Sqrt(9.0))
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
}
