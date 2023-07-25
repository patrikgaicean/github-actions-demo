package main

import "testing"

func TestTwoSum(t *testing.T) {
	got := twoSum(1, 2)
	if got != 3 {
		t.Errorf("twoSum(1, 2) = %d, want 3", got)
	}
}
