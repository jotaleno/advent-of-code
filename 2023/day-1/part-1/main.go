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
	totalScore := 0

	lines := strings.Split(puzzle, "\n")

	for _, line := range lines {
		totalScore += score(line)
	}

	return totalScore
}

func score(input string) int {
	var calibrationValue int
	var firstByte byte
	var lastByte byte

	inputLen := len(input)

	for i := 0; firstByte == 0 || lastByte == 0; i++ {
		fByte := input[i]
		lByte := input[inputLen-i-1]

		if firstByte == 0 && isByteANumber(fByte) {
			firstByte = fByte
		}
		if lastByte == 0 && isByteANumber(lByte) {
			lastByte = lByte
		}
	}

	code := fmt.Sprintf("%c%c", firstByte, lastByte)

	codeInt, err := strconv.Atoi(code)

	if err != nil {
		fmt.Printf("Could not convert string to int: %s\n", code)
	} else {
		calibrationValue = codeInt
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
