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
		totalScore += score(line)
	}

	return totalScore
}

func score(line string) int {
	var points int
	var winningNumbersMap = make(map[string]int)

	info := strings.Split(line, ":")

	list := strings.Split(info[1], "|")

	winningNumbers := list[0]
	cardNumbers := list[1]

	for _, w := range strings.Split(winningNumbers, " ") {
		if (len(w)) > 0 {
			trimW := strings.Trim(w, " ")

			winningNumbersMap[trimW] = 1
		}
	}

	for _, c := range strings.Split(cardNumbers, " ") {
		if (len(c)) > 0 {
			trimC := strings.Trim(c, " ")

			_, ok := winningNumbersMap[trimC]

			if ok {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
	}

	return points
}
