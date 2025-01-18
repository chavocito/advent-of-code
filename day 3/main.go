package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var aocInput string

func accumulateMuls(inp string) int {
	count := 0
	pattern := regexp.MustCompile("(mul\\(\\d+,\\d+\\))")
	matches := pattern.FindAll([]byte(inp), -1)
	for _, match := range matches {
		intArr, err := getInputIntegers(string(match))
		if err != nil {
			fmt.Println("Error getting integers from input:", err)
			continue
		}
		if len(intArr) == 0 {
			continue
		}
		count += executeMul(intArr)
	}
	return count
}

func getInputIntegers(input string) ([]int, error) {
	var outputArr []int
	input = strings.TrimPrefix(input, "mul(")
	input = strings.TrimSuffix(input, ")")
	arrStr := strings.Split(input, ",")
	for _, str := range arrStr {
		cleanInput, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		outputArr = append(outputArr, cleanInput)
	}
	return outputArr, nil
}

func executeMul(intInput []int) int {
	return intInput[0] * intInput[1]
}

func main() {
	count := accumulateMuls(aocInput)
	fmt.Println("Number of muls:", count)
}
