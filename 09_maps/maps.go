package main

import "fmt"

func main() {
	// ==========================================
	// 1. DECLARING AND INITIALIZING MAPS
	// ==========================================
	fmt.Println("--- 1. Initialization ---")

	// Method A: Using a map literal
	// Syntax: map[KeyType]ValueType{...}
	productPrices := map[string]float64{
		"laptop": 999.99,
		"phone":  599.50,
	}
	fmt.Println("Literal Map:", productPrices)

	// Method B: Creating an empty initialized map using 'make'
	// Use this when you don't know the values ahead of time.
	employeeIds := make(map[string]int)
	
	// Adding values to the map
	employeeIds["Alice"] = 1001
	employeeIds["Bob"] = 1002
	
	fmt.Println("Make Map:", employeeIds)

	// ==========================================
	// 2. UPDATING AND RETRIEVING VALUES
	// ==========================================
	fmt.Println("\n--- 2. Update & Retrieval ---")

	// Updating an existing key
	employeeIds["Alice"] = 5555 
	fmt.Println("Updated Alice ID:", employeeIds["Alice"])

	// Reading a key that DOES NOT exist
	// Go safely returns the default "zero value" for the data type (0 for int, "" for string)
	fmt.Println("Missing key returns zero value:", employeeIds["Charlie"])

	// ==========================================
	// 3. CHECKING IF A KEY EXISTS (COMMA OK)
	// ==========================================
	// Because missing keys return zero values, use the 'comma ok' idiom 
	// to verify if a key is genuinely present in the map.
	fmt.Println("\n--- 3. Checking Existence ---")

	// 'value' gets the result, 'ok' gets a boolean (true if key exists, false if missing)
	value, ok := employeeIds["Charlie"]
	if ok {
		fmt.Printf("Charlie found! ID is %d\n", value)
	} else {
		fmt.Println("Charlie does not exist in the map.")
	}

	// Shorthand syntax often used inside if statements:
	if id, exists := employeeIds["Bob"]; exists {
		fmt.Printf("Bob exists with ID: %d\n", id)
	}

	// ==========================================
	// 4. DELETING KEYS
	// ==========================================
	// Use the built-in delete() function.
	// Syntax: delete(mapName, key)
	fmt.Println("\n--- 4. Deleting Keys ---")
	
	fmt.Println("Before delete:", employeeIds)
	
	delete(employeeIds, "Bob") // Removes Bob from the map completely
	
	fmt.Println("After delete:", employeeIds)
	// Note: Deleting a key that doesn't exist does nothing and will not cause an error.

	// ==========================================
	// 5. ITERATING OVER A MAP
	// ==========================================
	// Loops over all keys and values using 'range'.
	// CRITICAL: The order of results changes randomly across program executions!
	fmt.Println("\n--- 5. Iterating Maps ---")

	inventory := map[string]int{
		"apples":  50,
		"oranges": 23,
		"bananas": 11,
	}

	for item, count := range inventory {
		fmt.Printf("We have %d %s left in stock\n", count, item)
	}
}
