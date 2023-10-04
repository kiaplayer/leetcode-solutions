package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMain_EqualFrequency(t *testing.T) {
	// arrange
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "0",
			expected: 0,
		},
		{
			input:    "42",
			expected: 42,
		},
		{
			input:    "   -42",
			expected: -42,
		},
		{
			input:    "4193 with words",
			expected: 4193,
		},
		{
			input:    "words and 987",
			expected: 0,
		},
		{
			input:    "-91283472332",
			expected: -2147483648,
		},
		{
			input:    "0032",
			expected: 32,
		},
		{
			input:    "+-12",
			expected: 0,
		},
		{
			input:    "00000-42a1234",
			expected: 0,
		},
		{
			input:    "20000000000000000000",
			expected: 2147483647,
		},
		{
			input:    "  0000000000012345678",
			expected: 12345678,
		},
		{
			input:    "2147483646",
			expected: 2147483646,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	for _, tt := range testCases {
		t.Run(tt.input, func(t *testing.T) {
			// act
			actual := MyAtoi(tt.input)
			// assert
			assert.Equal(t, tt.expected, actual)
		})
	}
}
