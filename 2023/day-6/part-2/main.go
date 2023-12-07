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
	lines := strings.Split(puzzle, "\n")

	times := lines[0][len("time:"):]
	distances := lines[1][len("distance:"):]

	var timesStr string
	for _, t := range strings.Fields(times) {
		timesStr = fmt.Sprintf("%s%s", timesStr, t)
	}

	var distancesStr string
	for _, d := range strings.Fields(distances) {
		distancesStr = fmt.Sprintf("%s%s", distancesStr, d)
	}

	time, _ := strconv.Atoi(timesStr)
	distance, _ := strconv.Atoi(distancesStr)

	totalScore := score(time, distance)

	return totalScore
}

func score(time int, distance int) int {
	var count int

	for charge := 0; charge < time; charge++ {
		if (time-charge)*charge > distance {
			count++
		}
	}

	return count
}
