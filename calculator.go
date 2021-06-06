// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

func colour() {
	var colour string = "blå"

	if colour == "blå" {
		var colour string = "röd"
		fmt.Println("färgen är %s", colour)
		// färgen är röd
	}
	fmt.Println("färgen är %s", colour)
	// färgen är blå igen..
}

// Multiply stuff with other stuff.
func Multiply(a, b int32) int32 {
	return a * b
	// return 1
}

// Divide stuff
func Divide(a, b float64) (float64, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("!#Div/0")
	}
	return a / b, nil
}
