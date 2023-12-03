package main

import (
	"fmt"
	"os"
	"strings"
)

const FILE_NAME = "input.txt"

type index struct {
	y int
	x int
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

	fmt.Println(puzzle)

	lines := strings.Split(puzzle, "\n")
	length := len(lines)

	for y := 0; y < length; y++ {
		for x := 0; x < length; x++ {
			if isAsterisc(lines[y][x]) {
				if isAdjacentToTwoNumbers(lines, y, x, length, length) {

				}
			}
		}
	}

	return totalScore
}

func getNumber(line string, index index) string {
	var number string

	return number
}
func isAdjacentToTwoNumbers(puzzle []string, y int, x int, yLen int, xLen int) bool {
	var adjCounter int
	var indexes []index

	// up
	if y-1 >= 0 {
		if isNumber(puzzle[y-1][x]) {
			indexes = append(indexes, index{y: y - 1, x: x})
			adjCounter++
		}
	}
	// up right
	if x+1 < xLen && y-1 >= 0 {
		if isNumber(puzzle[y-1][x+1]) {
			indexes = append(indexes, index{y: y - 1, x: x + 1})
			adjCounter++
		}
	}

	// right
	if x+1 < xLen {
		if isNumber(puzzle[y][x+1]) {
			indexes = append(indexes, index{y: y, x: x + 1})
			adjCounter++
		}
	}

	// down right
	if x+1 < xLen && y+1 < yLen {
		if isNumber(puzzle[y+1][x+1]) {
			indexes = append(indexes, index{y: y + 1, x: x + 1})
			adjCounter++
		}
	}

	// down
	if y+1 < yLen {
		if isNumber(puzzle[y+1][x]) {
			indexes = append(indexes, index{y: y + 1, x: x})
			adjCounter++
		}
	}

	// down left
	if x-1 >= 0 && y+1 < yLen {
		if isNumber(puzzle[y+1][x-1]) {
			indexes = append(indexes, index{y: y + 1, x: x + 1})
			adjCounter++
		}
	}

	// left
	if x-1 >= 0 {
		if isNumber(puzzle[y][x-1]) {
			indexes = append(indexes, index{y: y, x: x - 1})
			adjCounter++
		}
	}

	// up left
	if x-1 >= 0 && y-1 >= 0 {
		if isNumber(puzzle[y-1][x-1]) {
			indexes = append(indexes, index{y: y - 1, x: x - 1})
			adjCounter++
		}
	}

	fmt.Println(adjCounter)

	if adjCounter == 2 {
		fmt.Println(indexes)
		return true
	} else {
		return false
	}
}

func isNumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	} else {
		return false
	}
}

func isAsterisc(char byte) bool {
	if char == '*' {
		return true
	} else {
		return false
	}
}
