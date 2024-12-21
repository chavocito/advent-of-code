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
	fmt.Println("Safe Combinations: ", input)
}

func isSafe(levels []int) bool {
	for i := 0; i < len(levels)-1; i++ {
		adjacencyDiff := math.Abs(float64(levels[i] - levels[i+1]))
		if adjacencyDiff < 1 || adjacencyDiff > 3 {
			return false
		}
		if levels[i] == levels[i+1] {
			return false
		}
	}
	return true
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
		for _, part := range parts {
			vals, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil, err
			}
			collections = append(collections, []int{vals})
		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	return collections, nil
}
