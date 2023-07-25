package main

import "testing"

func TestTwoSum(t *testing.T) {
	got := twoSum(1, 2)
	if got != 3 {
		t.Errorf("twoSum(1, 2) = %d, want 3", got)
	}

	got2 := twoSum(2, 2)
	if got2 != 4 {
		t.Errorf("twoSum(2, 2) = %d, want 4", got2)
	}

}
