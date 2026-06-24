// Error handling is a very important concept in Go. Unlike languages that use exceptions
// (try/catch), Go handles errors by returning them as values.
package main

import (
	"fmt"
)


func getUser(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("invalid user id")
	}

	return "Moiz", nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
//  nill meaan no error
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}


func validateAge(age int) error {
	if age < 18 {
		return errors.New("age must be 18 or older")
	}
	return nil
}

func main() {
	err := validateAge(16)

	if err != nil {
		fmt.Println(err)
	}
}