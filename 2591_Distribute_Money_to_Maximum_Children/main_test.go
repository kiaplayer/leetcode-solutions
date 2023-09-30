package main

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMain_EqualFrequency(t *testing.T) {
	// arrange
	testCases := []struct {
		money    int
		children int
		expected int
	}{
		{
			money:    20,
			children: 3,
			expected: 1,
		},
		{
			money:    16,
			children: 2,
			expected: 2,
		},
		{
			money:    6,
			children: 2,
			expected: 0,
		},
		{
			money:    5,
			children: 6,
			expected: ResultNoWay,
		},
		{
			money:    4,
			children: 1,
			expected: ResultNoWay,
		},
		{
			money:    2,
			children: 2,
			expected: 0,
		},
		{
			money:    18,
			children: 2,
			expected: 1,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	for _, tt := range testCases {
		testName := fmt.Sprintf("Money: %d, childen: %d", tt.money, tt.children)
		t.Run(testName, func(t *testing.T) {
			// act
			actual := DistMoney(tt.money, tt.children)
			// assert
			assert.Equal(t, tt.expected, actual)
		})
	}
}
