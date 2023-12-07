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
	totalScore := 1

	lines := strings.Split(puzzle, "\n")

	times := lines[0][len("time:"):]
	distances := lines[1][len("distance:"):]

	var timesList []int
	for _, t := range strings.Fields(times) {
		n, _ := strconv.Atoi(t)

		timesList = append(timesList, n)
	}

	var distancesList []int
	for _, d := range strings.Fields(distances) {
		n, _ := strconv.Atoi(d)

		distancesList = append(distancesList, n)
	}

	for i := 0; i < len(timesList); i++ {
		totalScore *= score(timesList[i], distancesList[i])
	}

	return totalScore
}

func score(time int, distance int) int {
	var count int

	for charge := 1; charge < time; charge++ {
		if (time-charge)*charge > distance {
			count++
		}
	}

	return count
}
