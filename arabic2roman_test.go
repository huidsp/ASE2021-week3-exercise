package main

import "testing"

func TestA2R01(t *testing.T) {
	expected := "I"
	result := a2r(1)
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}

func TestA2R02(t *testing.T) {
	expected := "II"
	result := a2r(2)
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}
