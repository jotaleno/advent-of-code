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
			input:    "467..1141..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..",
			expected: 467835,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, totalScore(c.input))
	}
}
