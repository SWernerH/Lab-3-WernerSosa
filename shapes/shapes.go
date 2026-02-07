package main

import(
	"math"
)

//rectangle
type Rectangle struct{
	width float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

//circle
type Circle struct{
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

//triangle
type Triangle struct{
	base float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t Triangle) Perimeter() float64 {
	return 3 * t.base
}

//pointers

func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func (c *Circle) Scale(factor float64) {
	c.radius *= factor
}

func (t *Triangle) Scale(factor float64) {
	t.base *= factor
	t.height *= factor
}