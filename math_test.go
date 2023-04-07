package main

import "testing"

func TestSum(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Expected 3")
	}
}
