package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScore(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{
			input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
			expected: 6440,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, totalScore(c.input))
	}
}
