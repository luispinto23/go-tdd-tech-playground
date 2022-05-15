package calculator_test

import (
	"calculator"
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		a, b float64
		want float64
		err  error
	}{
		{a: 10, b: 5, want: 2, err: nil},
		{a: 120, b: 0, want: 0.0, err: errors.New("division by zero not allowed")},
		{a: 99, b: 0, want: 0.0, err: errors.New("division by zero not allowed")},
		{a: 9, b: 3, want: 3, err: nil},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		if err != nil {
			t.Fatalf("Divide(%f, %f): want no error for valid input, got %v", tc.a, tc.b, err)
		}
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
