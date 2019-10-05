package main

import (
	"testing"
)

func TestIsPositionAlreadyFilled(t *testing.T) {
	tests := []struct {
		grid         [3][3]string
		row          int
		col          int
		expectedBool bool
	}{
		{
			grid:         [3][3]string{{"X", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}},
			row:          0,
			col:          0,
			expectedBool: true,
		},
		{
			grid:         [3][3]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}},
			row:          0,
			col:          1,
			expectedBool: false,
		},
		{
			grid:         [3][3]string{{"-", "X", "-"}, {"O", "-", "-"}, {"L", "-", "-"}},
			row:          2,
			col:          0,
			expectedBool: false,
		},
	}
	for _, test := range tests {
		actual := isPositionAlreadyFilled(test.grid, test.row, test.col)
		if actual != test.expectedBool {
			t.Errorf("Actual %v, Expected %v", actual, test.expectedBool)
		}
	}
}

func TestAllPositionsFilled(t *testing.T) {
	tests := []struct {
		grid         [3][3]string
		expectedBool bool
	}{
		{
			grid:         [3][3]string{{"X", "O", "O"}, {"O", "X", "X"}, {"O", "X", "O"}},
			expectedBool: true,
		},
		{
			grid:         [3][3]string{{"X", "-", "O"}, {"O", "X", "X"}, {"O", "X", "X"}},
			expectedBool: false,
		},
	}
	for _, test := range tests {
		actual := allPositionsFilled(test.grid)
		if actual != test.expectedBool {
			t.Errorf("Actual %v, Expected %v", actual, test.expectedBool)
		}
	}
}

func TestHasAnyoneWonYet(t *testing.T) {
	tests := []struct {
		grid         [3][3]string
		expectedBool bool
	}{
		{
			grid:         [3][3]string{{"X", "O", "O"}, {"O", "X", "X"}, {"O", "X", "O"}},
			expectedBool: false,
		},
		{
			grid:         [3][3]string{{"X", "O", "O"}, {"O", "X", "X"}, {"O", "X", "X"}},
			expectedBool: true,
		},
		{
			grid:         [3][3]string{{"X", "X", "X"}, {"O", "X", "O"}, {"O", "X", "O"}},
			expectedBool: true,
		},
	}
	for _, test := range tests {
		actual := hasAnyoneWonYet(test.grid)
		if actual != test.expectedBool {
			t.Errorf("Actual %v, Expected %v", actual, test.expectedBool)
		}
	}
}

func TestCheckAllRows(t *testing.T) {
	tests := []struct {
		grid         [3][3]string
		expectedBool bool
	}{
		{
			grid:         [3][3]string{{"X", "X", "X"}, {"O", "X", "X"}, {"O", "X", "O"}},
			expectedBool: true,
		},
		{
			grid:         [3][3]string{{"X", "O", "O"}, {"O", "X", "X"}, {"O", "X", "X"}},
			expectedBool: false,
		},
		{
			grid:         [3][3]string{{"X", "X", "X"}, {"X", "X", "X"}, {"X", "X", "X"}},
			expectedBool: true,
		},
	}
	for _, test := range tests {
		actual := checkAllRows(test.grid)
		if actual != test.expectedBool {
			t.Errorf("Actual %v, Expected %v", actual, test.expectedBool)
		}
	}
}

func TestCheckAllColumns(t *testing.T) {
	tests := []struct {
		grid         [3][3]string
		expectedBool bool
	}{
		{
			grid:         [3][3]string{{"X", "O", "X"}, {"X", "X", "X"}, {"X", "O", "O"}},
			expectedBool: true,
		},
		{
			grid:         [3][3]string{{"X", "O", "O"}, {"O", "X", "X"}, {"O", "X", "X"}},
			expectedBool: false,
		},
		{
			grid:         [3][3]string{{"X", "X", "X"}, {"X", "X", "X"}, {"X", "X", "X"}},
			expectedBool: true,
		},
	}
	for _, test := range tests {
		actual := checkAllColumns(test.grid)
		if actual != test.expectedBool {
			t.Errorf("Actual %v, Expected %v", actual, test.expectedBool)
		}
	}
}

func TestCheckDiagonals(t *testing.T) {
	tests := []struct {
		grid         [3][3]string
		expectedBool bool
	}{
		{
			grid:         [3][3]string{{"X", "O", "X"}, {"X", "X", "X"}, {"X", "O", "X"}},
			expectedBool: true,
		},
		{
			grid:         [3][3]string{{"X", "O", "O"}, {"O", "X", "X"}, {"O", "X", "O"}},
			expectedBool: false,
		},
		{
			grid:         [3][3]string{{"X", "X", "X"}, {"X", "X", "X"}, {"X", "X", "X"}},
			expectedBool: true,
		},
	}
	for _, test := range tests {
		actual := checkDiagonals(test.grid)
		if actual != test.expectedBool {
			t.Errorf("Actual %v, Expected %v", actual, test.expectedBool)
		}
	}
}

func TestSameShape(t *testing.T) {
	tests := []struct {
		a            string
		b            string
		c            string
		expectedBool bool
	}{
		{
			a:            "X",
			b:            "X",
			c:            "X",
			expectedBool: true,
		},
		{
			a:            "X",
			b:            "O",
			c:            "X",
			expectedBool: false,
		},
		{
			a:            "O",
			b:            "O",
			c:            "O",
			expectedBool: true,
		},
	}
	for _, test := range tests {
		actual := sameShape(test.a, test.b, test.c)
		if actual != test.expectedBool {
			t.Errorf("Actual %v, Expected %v", actual, test.expectedBool)
		}
	}
}
