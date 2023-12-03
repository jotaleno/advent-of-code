package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILE_NAME = "input.txt"

type direction struct {
	dY int
	dX int
}

var directions = []direction{
	{
		dY: 1,
		dX: 0,
	}, {
		dY: 1,
		dX: 1,
	}, {
		dY: 0,
		dX: 1,
	}, {
		dY: 1,
		dX: 0,
	}, {
		dY: 1,
		dX: 0,
	}, {
		dY: 1,
		dX: 0,
	}, {
		dY: 1,
		dX: 0,
	}, {
		dY: 1,
		dX: 0,
	},
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
	length := len(lines)

	for y := 0; y < length; y++ {
		for x := 0; x < length; {

			str := getNumber(lines[y], x)

			if str != "" {
				xStart := x
				xEnd := xStart + len(str) - 1

				if isAdjacentToSymbols(lines, y, xStart, xEnd, length, length) {
					n, err := strconv.Atoi(str)

					if err != nil {
						fmt.Printf("Could not convert string to int: %s\n", str)
					} else {
						totalScore += n
					}
				}

				x += len(str)

			} else {
				x++
			}
		}
	}

	return totalScore
}

func getNumber(line string, x int) string {
	var number string

	curr := x

	for curr < len(line) && isNumber(line[curr]) {
		number = line[x : curr+1]
		curr++
	}

	return number
}

func isAdjacentToSymbols(puzzle []string, y int, xStart int, xEnd int, yLen int, xLen int) bool {
	for x := xStart; x <= xEnd; x++ {
		for j := -1; j <= 1; j++ {
			for i := -1; i <= 1; i++ {
				if j != 0 || i != 0 {
					if y+j >= 0 && y+j < yLen && x+i >= 0 && x+i < xLen {
						if !isSymbol(puzzle[y+j][x+i]) {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func isSymbol(char byte) bool {
	if !isNumber(char) && !isDot(char) {
		return false
	} else {
		return true
	}
}

func isNumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	} else {
		return false
	}
}

func isDot(char byte) bool {
	if char == '.' {
		return true
	} else {
		return false
	}
}
