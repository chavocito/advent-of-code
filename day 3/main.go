package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getMuls(inp string) {
	pattern := regexp.MustCompile("(mul\\(\\d+,\\d+\\))")
	matches := pattern.FindAll([]byte(inp), -1)
	for _, match := range matches {
		matchStr := string(match)
	}
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

func execMul(input string) int {
	x, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return x
}

func main() {
	getMuls(input)
}
