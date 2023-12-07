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
			input: `Time:      7  15   30
Distance:  9  40  200`,
			expected: 71503,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, totalScore(c.input))
	}
}
