package main

import (
	"math"
	"testing"
)

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// func TestPerimeter(t *testing.T) {
// 	got := Perimeter(10.0, 10.0)
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// func TestDifferentPerimeter(t *testing.T) {
// 	got := Perimeter(20.0, 10.0)
// 	want := 60.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

//////

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
	//return 72.0
}

func TestAreaofRectangle(t *testing.T) {

	got := Rectangle{Height: 12.0, Width: 6.0}.Area()
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//test for circle
func TestAreaOfCircle(t *testing.T) {
	circle := Circle{Radius: 10.0}
	got := circle.Area()
	want := 314.1592653589793

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

type Rectangle struct {
	Height float64
	Width  float64
}

func TestPerimeterofRectangle(t *testing.T) {
	rectangle := Rectangle{Height: 10.0, Width: 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func TestAreaOfShapes(t *testing.T) {
	cases := []struct {
		Description string
		Shape       Shape
		Want        float64
	}{
		{Description: "Circles", Shape: Circle{Radius: 10.0}, Want: 314.1592653589793},
		{Description: "Rectangles", Shape: Rectangle{Height: 10.0, Width: 10.0}, Want: 100},
		//todo: add triangle support
		{Description: "Triangle", Shape: Triangle{Base: 12.0, Height: 6.0}, Want: 36},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := test.Shape.Area()
			if got != test.Want {
				t.Errorf("got %.2f want %.2f", got, test.Want)
			}

		})
	}
}
