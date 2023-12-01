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
			input:    "two1nine",
			expected: 29,
		},
		{
			input:    "eightwothree",
			expected: 83,
		},
		{
			input:    "abcone2threexyz",
			expected: 13,
		},
		{
			input:    "xtwone3four",
			expected: 24,
		},
		{
			input:    "4nineeightseven2",
			expected: 42,
		},
		{
			input:    "zoneight234",
			expected: 14,
		},
		{
			input:    "7pqrstsixteen",
			expected: 76,
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
