package main

import (
	_ "embed"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func generateArrayFromInput() [][]int {
	var finalArr [][]int
	splitStr := strings.Split(input, "\n")
	for _, coord := range splitStr {
		append(finalArr, make([]int, strings.Split(coord, " ")))
	}
	return splitStr
}

func main() {
	generateArrayFromInput()
	//count := countSafePaths(input)
	//fmt.Println("Number of safe paths:", count)
}

func countSafePaths(array [][]int) int {
	var count int = 0
	for _, levels := range array {
		if safePostDampening(levels) {
			count++
		}
	}
	return count
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	increasing := levels[0] < levels[1]
	for i := 0; i < len(levels)-1; i++ {
		adjacencyDiff := math.Abs(float64(levels[i] - levels[i+1]))
		if adjacencyDiff < 1 || adjacencyDiff > 3 {
			return false
		}
		if (increasing && levels[i] >= levels[i+1]) || (!increasing && levels[i] <= levels[i+1]) {
			return false
		}
	}
	return true
}

func safePostDampening(levels []int) bool {
	increasing := levels[0] < levels[1]
	if isSafe(levels) {
		return true
	}

	for i := 0; i < len(levels)-1; i++ {
		adjacencyDiff := math.Abs(float64(levels[i] - levels[i+1]))
		if adjacencyDiff > 3 {
			//fmt.Println("Removing", append(levels[:i+1], levels[i+2:]...))
			return isSafe(append(levels[:i+1], levels[i+2:]...))
		}
		if (increasing && levels[i] >= levels[i+1]) || (!increasing && levels[i] <= levels[i+1]) {
			return isSafe(append(levels[:i], levels[i+1:]...))
		}
	}
	return false
}
