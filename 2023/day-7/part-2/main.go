package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const FILE_NAME = "input.txt"

var CARDS = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
var RANKS = []string{"fiveOfKind", "fourOfKind", "fullHouse", "threeOfKind", "twoPair", "onePair", "highCard"}

type win struct {
	hand string
	rank int
	bid  int
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
	var winners []win

	cardsRanking := createRankingMap[rune](CARDS)
	handRanking := createRankingMap[string](RANKS)

	lines := strings.Split(puzzle, "\n")

	for _, line := range lines {
		splitLine := strings.Fields(line)

		hand := splitLine[0]
		bid, _ := strconv.Atoi(splitLine[1])

		rank := scoreHand(hand, handRanking)

		winners = append(winners, win{hand: hand, bid: bid, rank: rank})
	}

	sort.SliceStable(winners, func(i, j int) bool {
		if winners[i].rank < winners[j].rank {
			return true
		} else if winners[i].rank > winners[j].rank {
			return false
		} else {
			for c := 0; c < len(winners[i].hand); c++ {
				if cardsRanking[rune(winners[i].hand[c])] < cardsRanking[rune(winners[j].hand[c])] {
					return true
				} else if cardsRanking[rune(winners[i].hand[c])] > cardsRanking[rune(winners[j].hand[c])] {
					return false
				} else {
					continue
				}
			}
			return false
		}
	})

	for index, win := range winners {
		totalScore += (index + 1) * win.bid
	}

	return totalScore
}

func createRankingMap[T comparable](list []T) map[T]int {
	ranking := make(map[T]int)

	for i, element := range list {
		rank := len(list) - i

		ranking[element] = rank
	}

	return ranking
}

func scoreHand(hand string, handRanking map[string]int) int {
	cardsCounter := make(map[rune]int)
	var jokersCounter int

	for _, card := range hand {
		if rune(card) == 'J' {
			jokersCounter++
		} else {
			cardsCounter[card] += 1
		}
	}

	for c := range cardsCounter {
		cardsCounter[c] += jokersCounter
	}

	uniqueCards := len(cardsCounter)

	var score int

	switch uniqueCards {
	case 0:
		score = handRanking["fiveOfKind"]

	case 1:
		score = handRanking["fiveOfKind"]

	case 2:
		for _, c := range cardsCounter {
			if c == 4 {
				score = handRanking["fourOfKind"]
				break
			} else {
				score = handRanking["fullHouse"]
			}
		}

	case 3:
		for _, c := range cardsCounter {
			if c == 3 {
				score = handRanking["threeOfKind"]
				break
			} else {
				score = handRanking["twoPair"]
			}
		}

	case 4:
		score = handRanking["onePair"]

	default:
		score = handRanking["highCard"]
	}

	return score
}
