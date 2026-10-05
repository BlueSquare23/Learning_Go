/*
    Example using a struct with receiver and pointer methods
*/
package main

import "fmt"

// 1. Define the struct (the data blueprint)
type Rectangle struct {
	width  float64
	height float64
}

// 2. Define a value receiver method (reads data, cannot modify the original struct)
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// 3. Define a pointer receiver method (can modify the original struct)
func (r *Rectangle) Scale(factor float64) {
	r.width = r.width * factor
	r.height = r.height * factor
}

func main() {
	// 4. Instantiate the struct (creates the "object")
	rect := Rectangle{width: 10, height: 5}

	// 5. Call the Area method
	fmt.Printf("Original Area: %.2f\n", rect.Area()) // Output: 50.00

	// 6. Call the Scale method (modifies rect in-place)
	rect.Scale(2)
	
	// 7. Check the updated values
	fmt.Printf("New Width: %.2f, New Height: %.2f\n", rect.width, rect.height) // Output: 20.00, 10.00
	fmt.Printf("New Area: %.2f\n", rect.Area())      // Output: 200.00
}

