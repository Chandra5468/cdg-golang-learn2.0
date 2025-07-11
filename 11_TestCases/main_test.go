package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("ADD(2,3) returned %d, expected %d", result, expected)
	}
}

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5}, {"negative numbers", -2, -3, -5},
		{"zero and positive", 0, 5, 5}, {"negative and positive", -4, 2, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := add(tt.a, tt.b)

			if result != tt.expected {
				t.Errorf("Add(%d,%d) returned %d but expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

/*
	Run all tests in the current package: go test
	Run specific test functions: go test -run TestAdd
	Run tests in verbose mode (showing detailed output): go test -v
*/
