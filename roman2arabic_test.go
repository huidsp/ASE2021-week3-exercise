package main

import "testing"

func TestR2A01(t *testing.T) {
	expected := 1
	result := r2a("I")
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}

func TestR2A02(t *testing.T) {
	expected := 2
	result := r2a("II")
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}

func TestR2A03(t *testing.T) {
	expected := []int{3, 4, 5, 6, 7, 8, 9}
	input := []string{"III", "IV", "V", "VI", "VII", "VIII", "IX"}
	for i, x := range input {
		if expected[i] != r2a(x) {
			t.Errorf("Expected %v, Result %v", expected[i], x)
		}
	}
}
