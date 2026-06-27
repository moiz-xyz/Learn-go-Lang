package main

import "fmt"

func main() {
	// ==========================================
	// 1. DECLARING AND INITIALIZING SLICES
	// ==========================================
	fmt.Println("--- 1. Declaration ---")

	// Method A: Using a slice literal (Note: No size is specified inside the brackets [])
	var numbers = []int{10, 20, 30}
	fmt.Printf("Literal Slice: %v, Length: %d, Capacity: %d\n", numbers, len(numbers), cap(numbers))

	// Method B: Creating an empty slice using 'make'
	// Syntax: make([]Type, length, capacity)
	// This allocates an underlying array with space to grow up to a capacity of 5.
	dynamicSlice := make([]int, 3, 5) // Starts with 3 elements (all initialized to 0)
	fmt.Printf("Make Slice: %v, Length: %d, Capacity: %d\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	// ==========================================
	// 2. MODIFYING AND ACCESSING ELEMENTS
	// ==========================================
	fmt.Println("\n--- 2. Accessing & Modifying ---")
	
	// Slices use 0-based indexing just like other languages
	numbers[1] = 99 // Changes 20 to 99
	fmt.Printf("Updated Slice: %v, Element at index 1: %d\n", numbers, numbers[1])

	// ==========================================
	// 3. GROWING A SLICE WITH APPEND
	// ==========================================
	fmt.Println("\n--- 3. Using Append ---")
	
	var list []string // Starts as a nil slice (Length: 0, Capacity: 0)
	
	// append() takes the original slice, adds the item, and returns the updated slice
	list = append(list, "apple")
	list = append(list, "banana", "cherry") // You can append multiple items at once
	
	fmt.Printf("After Append: %v, Length: %d, Capacity: %d\n", list, len(list), cap(list))

	// ==========================================
	// 4. SLICING AN EXISTING SLICE OR ARRAY
	// ==========================================
	// You can create a new slice by cutting a portion out of an existing slice/array.
	// Syntax: slice[start_index : end_index] (includes start, but excludes end)
	fmt.Println("\n--- 4. Slicing Operations ---")
	
	months := []string{"Jan", "Feb", "Mar", "Apr", "May"}
	
	q1 := months[0:3]     // Grabs index 0, 1, and 2 ("Jan", "Feb", "Mar")
	fromMarch := months[2:] // Omit end index to go all the way to the end
	upToApril := months[:4] // Omit start index to start from 0
	
	fmt.Println("Q1 Slice:", q1)
	fmt.Println("From March:", fromMarch)
	fmt.Println("Up to April:", upToApril)

	// CRITICAL WARNING: Slices share memory!
	// If you modify an element in a sub-slice, it changes the original slice too.
	q1[0] = "JANUARY"
	fmt.Println("Modified Q1:", q1)
	fmt.Println("Original Months (also changed!):", months)

	// ==========================================
	// 5. COPYING SLICES SAFELY
	// ==========================================
	// To safely duplicate a slice without sharing memory, use the built-in copy() function.
	fmt.Println("\n--- 5. Copying Slices Safely ---")
	
	original := []int{1, 2, 3}
	
	// You MUST initialize the destination slice with the correct length before copying
	duplicate := make([]int, len(original)) 
	
	copy(duplicate, original) // Copies data from 'original' into 'duplicate'
	
	duplicate[0] = 999 // Changing the copy will NOT affect the original slice
	fmt.Println("Original Slice:", original)
	fmt.Println("Duplicate Slice:", duplicate)
}
