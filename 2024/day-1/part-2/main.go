package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILE_NAME = "input.txt"

func main() {
	f, err := os.ReadFile(FILE_NAME)

	if err != nil {
		fmt.Printf("Could not open file: %s\n", FILE_NAME)

		os.Exit(-1)
	}

	points := totalScore(string(f))

	fmt.Println(points)
}

func totalScore(puzzle string) int {
	totalScore := 0
	leftMap := make(map[int]int)
	rightMap := make(map[int]int)

	lines := strings.Split(strings.TrimSpace(puzzle), "\n")

	for _, line := range lines {
		numbers := strings.Fields(line)

		leftNumber, _ := strconv.Atoi(numbers[0])
		rightNumber, _ := strconv.Atoi(numbers[1])

		if _, ok := leftMap[leftNumber]; ok {
			leftMap[leftNumber]++
		} else {
			leftMap[leftNumber] = 1
		}

		if _, ok := rightMap[rightNumber]; ok {
			rightMap[rightNumber]++
		} else {
			rightMap[rightNumber] = 1
		}
	}

	for leftValue, leftCount := range leftMap {
		if rightValue, ok := rightMap[leftValue]; ok {
			totalScore += leftValue * leftCount * rightValue
		}
	}

	return totalScore
}
