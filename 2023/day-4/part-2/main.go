package main

import (
	"fmt"
	"os"
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
	var totalScore int

	lines := strings.Split(puzzle, "\n")

	cardCounter := make(map[int]int)

	for index, line := range lines {
		winningNumbers, cardNumbers := splitLine(line)

		id := index + 1

		cardCounter[id]++

		winningsCount := findWinningsCount(winningNumbers, cardNumbers)

		for w := 1; w <= winningsCount; w++ {
			cardCounter[id+w] += cardCounter[id]
		}
	}

	for _, count := range cardCounter {
		totalScore += count
	}

	return totalScore
}

func splitLine(line string) (string, string) {
	info := strings.Split(line, ":")
	list := strings.Split(info[1], "|")

	winningNumbers := list[0]
	cardNumbers := list[1]

	return winningNumbers, cardNumbers
}

func findWinningsCount(winningNumbers string, cardNumbers string) int {
	var count int
	var winningNumbersMap = make(map[string]int)

	for _, w := range strings.Fields(winningNumbers) {
		winningNumbersMap[w] = 1
	}

	for _, c := range strings.Fields(cardNumbers) {
		_, ok := winningNumbersMap[c]

		if ok {
			count++
		}
	}

	return count
}
