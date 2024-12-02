package main

import (
	"fmt"
	"math"
	"os"
	"sort"
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
	leftList := []int{}
	rightList := []int{}

	lines := strings.Split(strings.TrimSpace(puzzle), "\n")

	for _, line := range lines {
		numbers := strings.Fields(line)

		leftNumber, _ := strconv.Atoi(numbers[0])
		rightNumber, _ := strconv.Atoi(numbers[1])

		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	for i := 0; i < len(leftList); i++ {
		distance := int(math.Abs(float64(leftList[i]) - float64(rightList[i])))

		totalScore += distance
	}

	return totalScore
}
