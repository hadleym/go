package main
import "testing"

func TestOdd(t *testing.T) {
	if !odd(3) {
		t.Errorf("3 was supposed to be odd")
	}

	if odd(4) {
		t.Errorf("4 was not supposed to be odd")
	}
}

func TestEven(t *testing.T) {
	if !even(2) {
		t.Errorf("2 was supposed to be even")
	}
}