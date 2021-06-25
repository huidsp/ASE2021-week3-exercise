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

func TestR2A04(t *testing.T) {
	expected := 10
	result := r2a("X")
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}

func TestR2A05(t *testing.T) {
	expected := []int{20, 30, 40, 50, 60, 70, 80, 90}
	input := []string{"XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	for i, x := range input {
		if expected[i] != r2a(x) {
			t.Errorf("Expected %v, Result %v", expected[i], x)
		}
	}
}

func TestR2A06(t *testing.T) {
	expected := 11
	result := r2a("XI")
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}

func TestR2A07(t *testing.T) {
	expected := 21
	result := r2a("XXI")
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}

func TestR2A08(t *testing.T) {
	expected := 34
	result := r2a("XXXIV")
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}

func TestR2A09(t *testing.T) {
	expected := 0
	result := r2a("IVXXX")
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}

func TestR2A10(t *testing.T) {
	expected := 434
	result := r2a("CDXXXIV")
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}

func TestR2A11(t *testing.T) {
	expected := 934
	result := r2a("CMXXXIV")
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}

func TestR2A12(t *testing.T) {
	expected := 2934
	result := r2a("MMCMXXXIV")
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}
