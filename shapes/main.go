package main

import(
	"fmt"
)

func main() {
	rect := Rectangle{width: 4, height: 5}
	circle := Circle{radius: 3}
	triangle := Triangle{base: 6, height: 4}

	fmt.Println("=== Shape Calculations ===")

	fmt.Printf("Rectangle Area: %.2f\n", rect.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", rect.Perimeter())

	fmt.Printf("Circle Area: %.2f\n", circle.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", circle.Perimeter())

	fmt.Printf("Triangle Area: %.2f\n", triangle.Area())
	fmt.Printf("Triangle Perimeter: %.2f\n", triangle.Perimeter())

	fmt.Println("\nScaling Rectangle by factor of 2...")
	rect.Scale(2)
	fmt.Printf("Scaled Rectangle Area: %.2f\n", rect.Area())
	fmt.Printf("Scaled Rectangle Perimeter: %.2f\n", rect.Perimeter())
}
