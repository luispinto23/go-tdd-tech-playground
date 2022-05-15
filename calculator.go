package calculator

import (
	"errors"
	"math"
)

// Divide returns the quotient of a and b.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

// Add returns the sum of two numbers.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference of two numbers.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Sqrt returns the square root of a number or error if the number is negative.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative number not allowed")
	}
	return math.Sqrt(a), nil
}
