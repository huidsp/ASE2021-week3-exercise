package main

import "testing"

func TestR2A01(t *testing.T) {
	expected := 1
	result := r2a("I")
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}
