package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILE_NAME = "input.txt"

var colorCount = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type cubes struct {
	color  string
	number int
}

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
		if len(line) > 0 {
			totalScore += score(line)
		}
	}

	return totalScore
}

func score(input string) int {
	minMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	info := strings.Split(input, ": ")
	gameList := strings.Split(info[1], "; ")

	for _, subset := range gameList {
		cbs := getCubes(subset)

		for _, c := range cbs {
			if c.number > minMap[c.color] {
				minMap[c.color] = c.number
			}
		}
	}

	return minMap["red"] * minMap["green"] * minMap["blue"]
}

func getCubes(s string) []cubes {
	var listOfCubes []cubes

	splitCubes := strings.Split(s, ", ")

	for _, cbs := range splitCubes {
		c := strings.Split(cbs, " ")

		numberAsStr := c[0]
		color := c[1]

		number, err := strconv.Atoi(numberAsStr)

		if err != nil {
			fmt.Printf("Could not convert %s to int\n", numberAsStr)
			os.Exit(-1)
		}

		listOfCubes = append(listOfCubes, cubes{
			color:  color,
			number: number,
		})
	}

	return listOfCubes
}
