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
		expected bool
	}{
		{
			input:    "aazz",
			expected: false,
		},
		{
			input:    "aabc",
			expected: true,
		},
		{
			input:    "bac",
			expected: true,
		},
		{
			input:    "cccaa",
			expected: true,
		},
		{
			input:    "abbcc",
			expected: true,
		},
		{
			input:    "cccd",
			expected: true,
		},
		{
			input:    "cbccca",
			expected: false,
		},
		{
			input:    "babbdd",
			expected: false,
		},
		{
			input:    "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz",
			expected: true,
		},
		{
			input:    "szzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz",
			expected: true,
		},
		{
			input:    "abcdefghijklmnopqrstuvwxyznabcdefghijklmnopqrstuvwxyz",
			expected: true,
		},
		{
			input:    "aaaabbbbccc",
			expected: false,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	for _, tt := range testCases {
		t.Run(tt.input, func(t *testing.T) {
			// act
			actual := EqualFrequency(tt.input)
			// assert
			assert.Equal(t, tt.expected, actual)
		})
	}
}
