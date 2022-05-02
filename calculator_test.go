package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	var want float64 = 15

	got := calculator.Add(5, 10)

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
