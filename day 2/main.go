package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const input = "input.txt"

func main() {
	input, err := separateInputs(input)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	count := countSafePaths(input)
	fmt.Println("Number of safe paths:", count)
}

func countSafePaths(array [][]int) int {
	var count int = 0
	for _, levels := range array {
		if isSafe(levels) {
			count++
		}
	}
	return count
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	levels, increasing := unsafeDampener(levels)
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

func unsafeDampener(levels []int) ([]int, bool) {
	increasing := levels[0] < levels[1]
	for i := 0; i < len(levels)-1; i++ {
		if (increasing && levels[i] >= levels[i+1]) || (!increasing && levels[i] <= levels[i+1]) {
			levels = append(levels[:i], levels[i+1:]...)
			break
		}
	}
	return levels, increasing
}

func separateInputs(filePath string) ([][]int, error) {
	var err error
	var collections [][]int

	//read input file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file ", filePath, " :", err)
		return nil, err
	}

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		var nums []int
		for _, part := range parts {
			vals, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil, err
			}
			nums = append(nums, vals)
		}
		collections = append(collections, nums)
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	return collections, nil
}
