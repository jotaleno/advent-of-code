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
			input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			expected: 31,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, totalScore(c.input))
	}
}
