package main

import(
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	const epsilon = 0.0001
	return math.Abs(a-b) < epsilon
}

// Area Tests

func TestRectangleArea(t *testing.T) {
	r := Rectangle{width: 4, height: 5}
	expected := 20.0
	if !almostEqual(r.Area(), expected) {
		t.Errorf("Rectangle Area = %v, want %v", r.Area(), expected)
	}
}

func TestCircleArea(t *testing.T) {
	c := Circle{radius: 3}
	expected := math.Pi * 3 * 3
	if !almostEqual(c.Area(), expected) {
		t.Errorf("Circle Area = %v, want %v", c.Area(), expected)
	}
}

func TestTriangleArea(t *testing.T) {
	tg := Triangle{base: 6, height: 4}
	expected := 0.5 * 6 * 4
	if !almostEqual(tg.Area(), expected) {
		t.Errorf("Triangle Area = %v, want %v", tg.Area(), expected)
	}
}

// Perimeter Tests

func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{width: 4, height: 5}
	expected := 18.0
	if !almostEqual(r.Perimeter(), expected) {
		t.Errorf("Rectangle Perimeter = %v, want %v", r.Perimeter(), expected)
	}
}

func TestCirclePerimeter(t *testing.T) {
	c := Circle{radius: 3}
	expected := 2 * math.Pi * 3
	if !almostEqual(c.Perimeter(), expected) {
		t.Errorf("Circle Perimeter = %v, want %v", c.Perimeter(), expected)
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tg := Triangle{base: 6, height: 4}
	expected := 18.0
	if !almostEqual(tg.Perimeter(), expected) {
		t.Errorf("Triangle Perimeter = %v, want %v", tg.Perimeter(), expected)
	}
}

// Tests

func TestRectangleScale(t *testing.T) {
	r := Rectangle{width: 4, height: 5}
	r.Scale(2)
	if !almostEqual(r.width, 8) || !almostEqual(r.height, 10) {
		t.Errorf("Rectangle Scale failed, got Width=%v Height=%v", r.width, r.height)
	}
}

func TestCircleScale(t *testing.T) {
	c := Circle{radius: 3}
	c.Scale(2)
	if !almostEqual(c.radius, 6) {
		t.Errorf("Circle Scale failed, got Radius=%v", c.radius)
	}
}

func TestTriangleScale(t *testing.T) {
	tg := Triangle{base: 6, height: 4}
	tg.Scale(2)
	if !almostEqual(tg.base, 12) || !almostEqual(tg.height, 8) {
		t.Errorf("Triangle Scale failed, got Base=%v Height=%v", tg.base, tg.height)
	}
}