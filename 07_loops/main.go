package main

import "fmt"

func main() {
	// ==========================================
	// 1. TRADITIONAL THREE-COMPONENT LOOP
	// ==========================================
	// Works like a standard loop in Java/C++/JavaScript.
	// Structure: for [initialization]; [condition]; [post statement]
	fmt.Println("--- 1. Traditional Loop ---")
	for i := 0; i < 3; i++ {
		fmt.Printf("Iteration number: %d\n", i)
	}

	// ==========================================
	// 2. THE "WHILE" LOOP EQUIVALENT
	// ==========================================
	// Go doesn't have a 'while' keyword. 
	// Dropping the initialization and post statements turns 'for' into a while loop.
	fmt.Println("\n--- 2. While Loop Equivalent ---")
	count := 1
	for count <= 3 {
		fmt.Printf("Count is currently: %d\n", count)
		count++ // Increment to avoid an infinite loop
	}

	// ==========================================
	// 3. INFINITE LOOP WITH BREAK
	// ==========================================
	// Providing no conditions creates a loop that runs forever.
	// We use 'break' inside to safely stop it when a condition is met.
	fmt.Println("\n--- 3. Infinite Loop with Break ---")
	cycles := 0
	for {
		cycles++
		if cycles > 2 {
			fmt.Println("Stopping the infinite loop safely using 'break'")
			break // Exits the loop immediately
		}
		fmt.Printf("Cycle: %d\n", cycles)
	}

	// ==========================================
	// 4. FOR-RANGE WITH SLICES (ARRAYS)
	// ==========================================
	// The 'range' keyword loops through collections.
	// It automatically returns two variables: the index and a copy of the value.
	fmt.Println("\n--- 4. For-Range with Slices ---")
	fruits := []string{"apple", "banana"}

	// Example A: Accessing both index and value
	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// Example B: Skipping the index using the blank identifier '_'
	// Use this when your code does not need to know the index position.
	for _, value := range fruits {
		fmt.Printf("Fruit name: %s\n", value)
	}

	// ==========================================
	// 5. FOR-RANGE WITH MAPS (KEY-VALUE)
	// ==========================================
	// Ranging over a map returns the key and its paired value.
	// Note: Go maps iterate in a random order every time you run the code.
	fmt.Println("\n--- 5. For-Range with Maps ---")
	userScores := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}
	for name, score := range userScores {
		fmt.Printf("%s achieved a score of %d\n", name, score)
	}

	// ==========================================
	// 6. LOOP CONTROL (CONTINUE)
	// ==========================================
	// 'continue' skips the current step and jumps straight to the next iteration.
	fmt.Println("\n--- 6. Loop Control with Continue ---")
	for number := 1; number <= 4; number++ {
		if number == 2 {
			fmt.Println("Skipping number 2 using 'continue'")
			continue // Skips the Printf below for this iteration
		}
		fmt.Printf("Processing number: %d\n", number)
	}
}
