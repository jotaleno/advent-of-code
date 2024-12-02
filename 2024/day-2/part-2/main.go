package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const FILE_NAME = "input.txt"

const MIN_DIFF = 1
const MAX_DIFF = 3

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

	lines := strings.Split(strings.TrimSpace(puzzle), "\n")

	for _, line := range lines {
		levels := strings.Fields(line)

		if isSafe(levels) {
			totalScore++
		} else {
			for i := 0; i < len(levels); i++ {
				tmp := append([]string(nil), levels...)

				if isSafe(removeIndex(tmp, i)) {
					totalScore++
					break
				}
			}
		}
	}

	return totalScore
}

func isSafe(report []string) bool {
	for i := 1; i < len(report)-1; i++ {
		previousNumber, _ := strconv.Atoi(report[i-1])
		currentNumber, _ := strconv.Atoi(report[i])
		nextNumber, _ := strconv.Atoi(report[i+1])

		previousDiff := int(math.Abs(float64(currentNumber - previousNumber)))
		nextDiff := int(math.Abs(float64(nextNumber - currentNumber)))

		if (previousDiff < MIN_DIFF || previousDiff > MAX_DIFF || nextDiff < MIN_DIFF || nextDiff > MAX_DIFF) || (currentNumber-previousNumber > 0 && nextNumber-currentNumber < 0) || (currentNumber-previousNumber < 0 && nextNumber-currentNumber > 0) {
			return false
		}
	}

	return true
}

func removeIndex(arr []string, index int) []string {
	return append(arr[:index], arr[index+1:]...)
}
