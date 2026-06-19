package main

import "fmt"

// 1. DEFINING THE STRUCT
// This acts as our blueprint. We define a custom type named 'Book'.
type Book struct {
	Title  string
	Author string
	Pages  int
	InStock bool
}

func main() {
	// 2. CREATING INSTANCES (How to initialize a struct)

	// Method A: Field-Value Initialization (Recommended)
	// You explicitly name the fields. Order doesn't matter here.
	book1 := Book{
		Title:   "The Go Programming Language",
		Author:  "Alan Donovan",
		Pages:   380,
		InStock: true,
	}

	// Method B: Value-Only Initialization
	// You omit the field names, but you MUST provide values in the EXACT order 
	// they are defined in the struct type.
	book2 := Book{"The Hobbit", "J.R.R. Tolkien", 310, false}

	// Method C: Zero-Value Initialization
	// If you don't assign values, Go gives every field its default "zero value"
	// (strings become "", ints become 0, booleans become false).
	var book3 Book


	// 3. USING THE STRUCT (Reading and writing data)

	// Reading data using dot notation (.)
	fmt.Printf("Book 1: '%s' by %s (%d pages)\n", book1.Title, book1.Author, book1.Pages)
	fmt.Println("Is Book 2 in stock?", book2.InStock)

	// Modifying (Writing) data using dot notation (.)
	fmt.Println("Initial Book 3 Title:", book3.Title) // Will be empty ""
	
	book3.Title = "Dune"
	book3.Author = "Frank Herbert"
	book3.Pages = 600
	book3.InStock = true

	fmt.Println("Updated Book 3 Title:", book3.Title) // Now prints "Dune"


	// 4. STRUCT AS FUNCTION ARGUMENTS
	// You can pass structs into functions just like any other variable.
	printBookDetails(book1)
	printBookDetails(book3)
}

// A function that accepts our custom 'Book' struct as a parameter
func printBookDetails(b Book) {
	status := "Available"
	if !b.InStock {
		status = "Out of Stock"
	}
	fmt.Printf("[%s] '%s' by %s\n", status, b.Title, b.Author)
}