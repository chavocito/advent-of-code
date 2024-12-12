package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const input = "input.txt"

func main() {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	leftSlice, rightSlice, err := separateInputs(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	slices.Sort(leftSlice)
	slices.Sort(rightSlice)

	fmt.Println("Total difference: ", accumulator(leftSlice, rightSlice))
	fmt.Println("Similarity Score: ", similarityScore(leftSlice, rightSlice))
}

func accumulator(left []int, right []int) int {
	total := 0
	for i := 0; i < len(left); i++ {
		diff := int(math.Abs(float64(left[i] - right[i])))
		total += diff
	}
	return total
}

func separateInputs(file *os.File) ([]int, []int, error) {
	// Create arrays to hold the values
	var err error
	var leftValues []int
	var rightValues []int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			left, err1 := strconv.Atoi(parts[0])
			right, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				leftValues = append(leftValues, left)
				rightValues = append(rightValues, right)
			}
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil, err
	}

	return leftValues, rightValues, nil
}

func similarityScore(left []int, right []int) int {
	score := 0
	for _, el := range left {
		score += el * frequency(el, right)
	}
	return score
}

func frequency(el int, arr []int) int {
	freq := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == el {
			freq++
		}
	}
	return freq
}
