package shapes

import "math"

type Shape interface {
	Area() float64
}

type Square struct {
	Width  float64
	Height float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius  float64
}

type Triangle struct {
	Base  float64
	Height  float64
}

func (s Square) Area() float64 {
	return s.Height * s.Width
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}

var areaTests = []struct {
	name string
	shape Shape
 	want float64
}{
	{"Rectangle", Rectangle{2, 4}, 8},
 	{"Circle", Circle{1}, 3.141592653589793},
  {"Triangle", Triangle{1, 1}, 0.5},
}