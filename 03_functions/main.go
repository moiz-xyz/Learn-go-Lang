package main

import (
	"errors"
	"fmt"
)

// 1. The Anatomy of a Basic Function
// A standard function definition in Go uses the 'func' keyword, followed by the function name, 
// parameters (with types), and the return type.
// Note: Since 'a' and 'b' are both integers, we can use the shorthand 'a, b int'.
func add(a, b int) int {
	return a + b
}

// 2. Returning Multiple Values
// Go functions can return more than one value. This is universally used in Go for error handling, 
// allowing a function to return both the desired result and an error status.
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		// If the divisor is 0, we return a default result (0) and an error object
		return 0, errors.New("cannot divide by zero")
	}
	// If everything is fine, we return the result and 'nil' for the error
	return dividend / divisor, nil
}

// 3. Named Return Values
// You can pre-declare the names of your return variables directly in the function signature.
// Go initializes them with their zero-values, and a "naked" return automatically sends them back.
func rectangleAreaAndPerimeter(width, height float64) (area float64, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // Automatically returns 'area' and 'perimeter'
}

// 4. Variadic Functions
// By using '...', a function can accept an arbitrary number of trailing arguments.
// Inside the function, 'numbers' behaves like a slice ([]int).
func sumAll(numbers ...int) int {
	total := 0
	// We use a blank identifier (_) because we only care about the value, not the loop index
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// --- Testing 1. Basic Function ---
	result := add(10, 2)
	fmt.Println("Add Result:", result) // Output: 12

	// --- Testing 2. Multiple Return Values ---
	// Case A: Successful division
	divResult, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide Result:", divResult) // Output: 5
	}

	// Case B: Division by zero error
	_, errZero := divide(10, 0)
	if errZero != nil {
		fmt.Println("Divide Error handled:", errZero) // Output: cannot divide by zero
	}

	// --- Testing 3. Named Return Values ---
	a, p := rectangleAreaAndPerimeter(5, 4)
	fmt.Printf("Rectangle -> Area: %v, Perimeter: %v\n", a, p) // Output: Area: 20, Perimeter: 18

	// --- Testing 4. Variadic Function ---
	totalSum := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Variadic Sum:", totalSum) // Output: 15

	// --- Testing 5. First-Class Functions (Anonymous/Variables) ---
	// Functions can be assigned to variables right inside your logic execution block.
	multiply := func(x, y int) int {
		return x * y
	}
	fmt.Println("Multiply Result:", multiply(3, 4)) // Output: 12
}