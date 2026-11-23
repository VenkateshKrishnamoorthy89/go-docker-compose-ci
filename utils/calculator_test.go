package utils

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 10, 20, 30},
		{"negative numbers", -10, -20, -30},
		{"mixed numbers", 10, -5, 5},
		{"zero", 0, 0, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)

			if result != tc.expected {
				t.Errorf(
					"Add(%d, %d) = %d; expected %d",
					tc.a, tc.b, result, tc.expected,
				)
			}
		})
	}
}
