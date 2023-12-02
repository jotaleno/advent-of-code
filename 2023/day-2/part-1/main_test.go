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
			input:    "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected: 1,
		},
		{
			input:    "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expected: 2,
		},
		{
			input:    "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expected: 0,
		},
		{
			input:    "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			expected: 0,
		},
		{
			input:    "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			expected: 5,
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.expected, score(c.input, i+1))
	}
}
