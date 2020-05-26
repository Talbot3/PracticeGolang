package mymath

import (
	"testing"
)

func TestGcd(t *testing.T) {
	valA := 22
	valB := 33
	result := Gcd(valA, valB)
	if result != 11 {
		t.Errorf("Gcd(%d, %d) = 11, but result %d", valA, valB, result)
	}
}

func TestFib(t *testing.T) {
	valA := 3
	result := Fib(valA)
	if result != 2 {
		t.Errorf("Fib(%d) = 3, but result %d", valA, result)
	}
}
