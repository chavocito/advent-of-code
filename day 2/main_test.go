package main

import "testing"

var mockInput = [][]int{
	{16, 10, 15, 5, 1},
	{14, 17, 20, 23, 25},
	{11, 11, 13, 14, 5},
	{21, 19, 17, 15},
}

func TestIsSafe(t *testing.T) {
	safeInputs := [][]int{
		{14, 17, 20, 23, 25},
		{21, 19, 17, 15},
	}
	for _, levels := range safeInputs {
		if !isSafe(levels) {
			t.Error("Expected true, got false")
		} else {
			t.Log("isSafe passed")
		}
	}
}

func TestCountSafePath(t *testing.T) {
	if countSafePaths(mockInput) == 2 {
		t.Log("countSafePaths passed")
		return
	}
	t.Error("Expected 2, got ", countSafePaths(mockInput))
}
