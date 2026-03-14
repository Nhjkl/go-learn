package main

import "testing"

func TestExample(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"case1", 1, 2},
		{"case2", 2, 4},
		{"case3", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Example(tt.input)
			if got != tt.expected {
				t.Errorf("Example(%d) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}
