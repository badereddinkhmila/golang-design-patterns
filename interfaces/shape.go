package interfaces

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

// does not implement Shape interface !!
type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (s Square) Area() float64 {
	return math.Pow(s.side, 2)
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (C Circle) Perimeter() float64 {
	return 2 * C.radius * math.Pi
}

func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func TotalPerimeter(shapes ...Shape) float64 {
	var peri float64
	for _, s := range shapes {
		peri += s.Perimeter()
	}
	return peri
}
