package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getCalculablesFromInput() [][]string {
	var inputPairs [][]string
	result := strings.Split(input, "mul(")
	for _, r := range result {
		insideSplit := strings.Split(r, ",")
		if len(insideSplit) == 2 {
			secondEl := strings.Split(insideSplit[1], ")")
			inputPairs = append(inputPairs, []string{insideSplit[0], secondEl[0]})
		}
	}
	return inputPairs
}

func main() {
	var count int = 0
	inputPairs := getCalculablesFromInput()
	for _, pair := range inputPairs {
		x, err := strconv.Atoi(pair[0])
		y, err := strconv.Atoi(pair[1])

		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}
		mul := x * y
		count += mul
	}
	fmt.Println(count)
}
