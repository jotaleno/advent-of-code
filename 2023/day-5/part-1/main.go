package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const SOURCE = "source"
const DESTINATION = "destination"

const FILE_NAME = "input.txt"

func main() {
	f, err := os.ReadFile(FILE_NAME)

	if err != nil {
		fmt.Printf("Could not open file: %s\n", FILE_NAME)

		os.Exit(-1)
	}

	points := score(string(f))

	fmt.Println(points)
}

func score(puzzle string) int {
	decoderMap := make(map[string]map[int]int)
	order := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	lines := strings.Split(puzzle, "\n")

	seeds := strings.Split(lines[0], "seeds: ")[1]

	for l := 2; l < len(lines); {

		fromTo := strings.Split(lines[l], " ")[0]
		decoderMap[fromTo] = make(map[int]int)

		var c int
		for c = l + 1; c < len(lines) && len(lines[c]) > 0 && lines[c][0] != '\r'; c++ {
			decoder := strings.Split(lines[c], " ")

			destination, _ := strconv.Atoi(decoder[0])
			source, _ := strconv.Atoi(decoder[1])
			range_, _ := strconv.Atoi(decoder[2])

			for r := 0; r < range_; r++ {
				decoderMap[fromTo][source+r] = destination + r
			}
		}

		l = c + 1
	}

	locations := []int{}

	for _, s := range strings.Split(seeds, " ") {
		seed, _ := strconv.Atoi(s)

		current := seed
		for _, o := range order {
			v, ok := decoderMap[o][current]

			if ok {
				current = v
			}
		}

		locations = append(locations, current)
	}

	minLocation := math.MaxInt

	for _, l := range locations {
		if l < minLocation {
			minLocation = l
		}
	}

	fmt.Println(locations)

	return minLocation
}
