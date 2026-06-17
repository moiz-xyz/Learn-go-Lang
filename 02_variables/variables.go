package main

import "fmt"

func main() {
	// -----------------------------------------------------------------
	// 1. TEXT TYPES
	// -----------------------------------------------------------------
	
	// string: Used for text. Must be enclosed in double quotes "".
	var username string = "Moiz"

	// -----------------------------------------------------------------
	// 2. WHOLE NUMBER TYPES (INTEGERS)
	// -----------------------------------------------------------------
	
	// int: The standard type for whole numbers (positive or negative).
	// On your Ubuntu 64-bit laptop, "int" automatically uses 64 bits of memory.
	var age int = 24

	// int8, int16, int32, int64: Integers with strict memory sizes.
	// int8 can only hold numbers from -128 to 127. Great for saving RAM.
	var smallNumber int8 = 100

	// uint: "Unsigned Integer". It can ONLY hold positive numbers (no minus sign).
	// Because it ignores negative numbers, it can hold much larger positive numbers.
	var positiveOnly uint = 550

	// -----------------------------------------------------------------
	// 3. DECIMAL NUMBER TYPES (FLOATING POINT)
	// -----------------------------------------------------------------
	
	// float32 & float64: Used for numbers with decimal points.
	// You must explicitly choose 32 or 64 bits. float64 is the standard and most precise.
	var price float64 = 19.99

	// -----------------------------------------------------------------
	// 4. LOGICAL TYPES
	// -----------------------------------------------------------------
	
	// bool: "Boolean". Can only be either "true" or "false".
	// Used constantly for making decisions (if/else statements) in code.
	var isCoding bool = true

	// -----------------------------------------------------------------
	// 5. THE SHORT CUT (TYPE INFERENCE)
	// -----------------------------------------------------------------
	
	// The ":=" operator is shorthand. You do not type "var" or the data type.
	// Go looks at the value "Ubuntu" and automatically decides this is a string.
	operatingSystem := "Ubuntu"

	// -----------------------------------------------------------------
	// PRINTING THE OUTPUT
	// -----------------------------------------------------------------
	fmt.Println("--- Text Type ---")
	fmt.Println("Username:", username)

	fmt.Println("\n--- Integer Types ---")
	fmt.Println("Age (int):", age)
	fmt.Println("Small Number (int8):", smallNumber)
	fmt.Println("Positive Only (uint):", positiveOnly)

	fmt.Println("\n--- Decimal Type ---")
	fmt.Println("Price (float64):", price)

	fmt.Println("\n--- Boolean Type ---")
	fmt.Println("Is Coding?:", isCoding)

	fmt.Println("\n--- Shorthand Type ---")
	fmt.Println("OS:", operatingSystem)
}
