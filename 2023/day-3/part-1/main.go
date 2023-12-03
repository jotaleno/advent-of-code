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
		number = fmt.Sprintf("%s%c", number, line[curr])
		curr++
	}

	return number
}

func isAdjacentToSymbols(puzzle []string, y int, xStart int, xEnd int, yLen int, xLen int) bool {
	for x := xStart; x <= xEnd; x++ {
		// up
		if y-1 >= 0 {
			if !isSymbol(puzzle[y-1][x]) {
				return true
			}
		}
		// up right
		if x+1 < xLen && y-1 >= 0 {
			if !isSymbol(puzzle[y-1][x+1]) {
				return true
			}
		}

		// right
		if x+1 < xLen {
			if !isSymbol(puzzle[y][x+1]) {
				return true
			}
		}

		// down right
		if x+1 < xLen && y+1 < yLen {
			if !isSymbol(puzzle[y+1][x+1]) {
				return true
			}
		}

		// down
		if y+1 < yLen {
			if !isSymbol(puzzle[y+1][x]) {
				return true
			}
		}

		// down left
		if x-1 >= 0 && y+1 < yLen {
			if !isSymbol(puzzle[y+1][x-1]) {
				return true
			}
		}

		// left
		if x-1 >= 0 {
			if !isSymbol(puzzle[y][x-1]) {
				return true
			}
		}

		// up left
		if x-1 >= 0 && y-1 >= 0 {
			if !isSymbol(puzzle[y-1][x-1]) {
				return true
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
