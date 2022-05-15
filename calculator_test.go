package calculator_test

import (
	"calculator"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	var want float64 = 15

	got := calculator.Add(10, 5)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		a, b float64
		want float64
		err  error
	}{
		{a: 10, b: 5, want: 2},
		{a: 120, b: 0, want: 0.0},
		{a: 9, b: 3, want: 3},
		{a: 99, b: 0, want: 0.0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Divide %f by %f", tc.a, tc.b), func(t *testing.T) {
			got, err := calculator.Divide(tc.a, tc.b)

			if err != nil {
				t.Fatalf("Divide(%f, %f): want no error for valid input, got %v", tc.a, tc.b, err)
			}
			if tc.want != got {
				t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
			}
		})
	}
}
