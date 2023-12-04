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

	for _, line := range lines {
		winningNumbers, cardNumbers := splitLine(line)
		totalScore += score(winningNumbers, cardNumbers)
	}

	return totalScore
}

func splitLine(line string) (string, string) {
	info := strings.Split(line, ":")
	list := strings.Split(info[1], "|")

	return list[0], list[1]
}

func score(winningNumbers string, cardNumbers string) int {
	var points int
	var winningNumbersMap = make(map[string]int)

	for _, w := range strings.Fields(winningNumbers) {
		winningNumbersMap[w] = 1
	}

	for _, c := range strings.Fields(cardNumbers) {
		_, ok := winningNumbersMap[c]

		if ok {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}

	return points
}
