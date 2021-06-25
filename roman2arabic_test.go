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
