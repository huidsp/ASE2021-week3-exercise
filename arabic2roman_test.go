package main

import "testing"

func TestA2R01(t *testing.T) {
	expected := "I"
	result := a2r(1)
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}
