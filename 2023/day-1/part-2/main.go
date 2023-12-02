package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILE_NAME = "input.txt"

var mapValues = map[string]byte{
	"one":   49,
	"two":   50,
	"three": 51,
	"four":  52,
	"five":  53,
	"six":   54,
	"seven": 55,
	"eight": 56,
	"nine":  57,
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
	totalScore := 0

	lines := strings.Split(puzzle, "\n")

	for _, line := range lines {
		totalScore += score(line)
	}

	return totalScore
}

func score(input string) int {
	var calibrationValue int

	inputLen := len(input)

	if inputLen > 0 {
		var firstByte byte

		var lastByte byte

		for i := 0; firstByte == 0 || lastByte == 0; i++ {
			fByte := input[i]
			lByte := input[inputLen-i-1]

			if firstByte == 0 {
				if isByteANumber(fByte) {
					firstByte = fByte
				} else {
					possibleKey := findPossibleKey(input, i, false)

					if possibleKey != "" {
						val := isStringANumber(possibleKey)

						if val > 0 {
							firstByte = val
						}
					}
				}
			}
			if lastByte == 0 {
				if isByteANumber(lByte) {
					lastByte = lByte
				} else {
					possibleKey := findPossibleKey(input, i, true)

					if possibleKey != "" {
						val := isStringANumber(possibleKey)

						if val > 0 {
							lastByte = val
						}
					}
				}
			}
		}

		code := fmt.Sprintf("%c%c", firstByte, lastByte)

		codeInt, err := strconv.Atoi(code)

		if err != nil {
			fmt.Printf("Could not convert string to int: %s\n", code)
		} else {
			calibrationValue = codeInt
		}
	}

	return calibrationValue
}

func isByteANumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	} else {
		return false
	}
}

func isStringANumber(str string) byte {
	value, ok := mapValues[str]

	if ok {
		return value
	} else {
		return 0
	}
}

func findPossibleKey(s string, start int, isReversed bool) string {
	var found string

	for key := range mapValues {
		currLen := 0

		if !isReversed {
			for i := 0; i < len(key); i++ {
				if s[start+i] == key[i] {
					currLen++
				} else {
					break
				}
			}

			if currLen == len(key) {
				found = s[start : start+currLen]
				break
			}
		} else {
			for i := 0; i < len(key); i++ {
				if s[len(s)-start-1-i] == key[len(key)-i-1] {
					currLen++
				} else {
					break
				}
			}

			if currLen == len(key) {
				found = s[len(s)-len(key)-start : len(s)-start]
				break
			}
		}
	}
	return found
}
