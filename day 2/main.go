package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func generateArrayFromInput() ([][]int, error) {
	var finalArr [][]int
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		pointsAsStrings := strings.Split(strings.TrimSpace(line), " ")
		var pointsArr []int
		for _, pointStr := range pointsAsStrings {
			if pointStr == "" {
				continue
			}
			ints, err := strconv.Atoi(strings.TrimSpace(pointStr))
			if err != nil {
				return nil, fmt.Errorf("invalid input: %w", err)
			}
			pointsArr = append(pointsArr, ints)
		}
		finalArr = append(finalArr, pointsArr)
	}
	return finalArr, nil
}

func main() {
	fromInput, err := generateArrayFromInput()
	if err != nil {
		return
	}
	count := countSafePaths(fromInput)
	fmt.Println("Number of safe paths:", count)
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
	if len(levels) < 2 {
		return false
	}

	if isSafe(levels) {
		return true
	}

	increasing := levels[0] < levels[1]
	for i := 0; i < len(levels)-1; i++ {
		adjacencyDiff := math.Abs(float64(levels[i] - levels[i+1]))
		//check for first instance of constraint violation
		if adjacencyDiff < 1 || adjacencyDiff > 3 {
			levels = append(levels[:i], levels[i+1:]...)
			//
			break
		}
		if (increasing && levels[i] >= levels[i+1]) || (!increasing && levels[i] <= levels[i+1]) {
			levels = append(levels[:i], levels[i+1:]...)
			//return isSafe(levels)
			break
		}
	}
	return isSafe(levels)
}
