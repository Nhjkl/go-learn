package main

import "testing"

func TestSum(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	got := Sum(arr)
	want := 15
	if got != want {
		t.Errorf("Sum(%v) = %d; want %d", arr, got, want)
	}
}

func TestSumEmpty(t *testing.T) {
	arr := [5]int{}
	got := Sum(arr)
	want := 0

	if got != want {
		t.Errorf("Sum(%v) = %d; want %d", arr, got, want)
	}
}
