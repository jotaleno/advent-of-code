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
			input:    "1abc2",
			expected: 12,
		},
		{
			input:    "pqr3stu8vwx",
			expected: 38,
		},
		{
			input:    "a1b2c3d4e5f",
			expected: 15,
		},
		{
			input:    "treb7uchet",
			expected: 77,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, score(c.input))
	}
}

func TestIsByteANumber(t *testing.T) {
	cases := []struct {
		input    byte
		expected bool
	}{
		{
			input:    48,
			expected: true,
		},
		{
			input:    57,
			expected: true,
		},
		{
			input:    47,
			expected: false,
		},
		{
			input:    58,
			expected: false,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, isByteANumber(c.input))
	}
}
