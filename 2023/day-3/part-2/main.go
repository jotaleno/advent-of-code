package main

import (
	"fmt"
	"os"
	"strconv"
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

	lines := strings.Split(puzzle, "\n")
	length := len(lines)

	for y := 0; y < length; y++ {
		for x := 0; x < length; x++ {
			if isAsterisc(lines[y][x]) {
				isAdj, indexes := isAdjacentToTwoNumbers(lines, y, x, length, length)

				if isAdj {
					multiplication := 1

					for _, index := range indexes {
						str, _ := getNumber(lines[index.y], index.x)

						n, err := strconv.Atoi(str)

						if err != nil {
							fmt.Printf("Could not convert string to int: %s\n", str)
						} else {
							multiplication *= n
						}
					}

					totalScore += multiplication
				}
			}
		}
	}

	return totalScore
}

func getNumber(line string, xStart int) (string, [2]int) {
	var leftX int
	var rightX int

	for leftX = xStart; leftX >= 0 && isNumber(line[leftX]); {
		leftX--
	}

	for rightX = xStart; rightX < len(line) && isNumber(line[rightX]); {
		rightX++
	}

	return line[leftX+1 : rightX], [2]int{leftX, rightX}
}

func isAdjacentToTwoNumbers(puzzle []string, y int, x int, yLen int, xLen int) (bool, []index) {
	var adjCounter int
	var indexes []index

	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if j != 0 || i != 0 {
				if y+j >= 0 && y+j < yLen && x+i >= 0 && x+i < xLen {
					if isNumber(puzzle[y+j][x+i]) {
						var found bool

						for _, index := range indexes {
							_, rg := getNumber(puzzle[index.y], index.x)

							if index.y == y+j && (x+i >= rg[0] && x+i <= rg[1]) {
								found = true
							}
						}

						if !found {
							indexes = append(indexes, index{y: y + j, x: x + i})
							adjCounter++
						}
					}
				}
			}
		}
	}

	if adjCounter == 2 {
		return true, indexes
	} else {
		return false, nil
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
