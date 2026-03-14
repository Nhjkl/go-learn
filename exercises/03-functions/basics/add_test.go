package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 1, 2, 3},
		{"negative", -1, -2, -3},
		{"zero", 0, 0, 0},
		{"mixed", -1, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(3, 4)
	want := 12
	if got != want {
		t.Errorf("Multiply(3, 4) = %d; want %d", got, want)
	}
}

func TestSwap(t *testing.T) {
	a, b := Swap("hello", "world")
	if a != "world" || b != "hello" {
		t.Errorf("Swap(hello, world) = (%s, %s); want (world, hello)", a, b)
	}
}
