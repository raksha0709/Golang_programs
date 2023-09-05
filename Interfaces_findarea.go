package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func PrintArea(s Shape) {
	fmt.Printf("The Area is %0.2f\n", s.Area())
}
func main() {
	var r float64
	fmt.Println("Enter the radius:")
	fmt.Scan(&r)
	circle := Circle{radius: r}
	var w, h float64
	fmt.Println("enter the width:")
	fmt.Scan(&w)
	fmt.Println("enter the height:")
	fmt.Scan(&h)
	rectangle := Rectangle{width: w, height: h}
	PrintArea(circle)
	PrintArea(rectangle)
}
