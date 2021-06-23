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

func TestA2R03(t *testing.T) {
	expected := "III"
	result := a2r(3)
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}
}

func TestA2R04(t *testing.T) {
	expected := []string{"IV", "V", "VI", "VII", "VIII", "IX"}
	input := []int{4, 5, 6, 7, 8, 9}
	for i, x := range expected {
		result := a2r(input[i])
		if x != result {
			t.Errorf("Expected %v, Result %v", expected[i], result)
		}
	}
}

func TestA2R05(t *testing.T) {
	expected := "X"
	result := a2r(10)
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}

}

func TestA2R06(t *testing.T) {
	expected := "XX"
	result := a2r(20)
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}

}

func TestA2R07(t *testing.T) {
	expected := []string{"XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	input := []int{30, 40, 50, 60, 70, 80, 90}
	for i, x := range expected {
		result := a2r(input[i])
		if x != result {
			t.Errorf("Expected %v, Result %v", expected[i], result)
		}
	}
}

func TestA2R08(t *testing.T) {
	expected := "XI"
	result := a2r(11)
	if expected != result {
		t.Errorf("Expected %v, Result %v", expected, result)
	}

}
