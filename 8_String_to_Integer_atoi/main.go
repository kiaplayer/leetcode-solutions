package main

import (
	"math"
	"strings"
	"unicode"
)

const (
	SignUnknown  = 0
	SignPositive = 1
	SignNegative = -1
	MinInt32     = "2147483648"
	MaxInt32     = "2147483647"
)

var digitsMap = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func main() {
	// Run tests with "go test ."
}

func MyAtoi(s string) int {
	sign := SignUnknown
	digitsSubstring := ""
	for _, r := range s {
		if unicode.IsSpace(r) {
			if sign == SignUnknown {
				continue
			} else {
				break
			}
		} else if unicode.IsDigit(r) {
			if sign == SignUnknown {
				sign = SignPositive
			}
			digitsSubstring = digitsSubstring + string(r)
		} else if sign == SignUnknown && (r == '+' || r == '-') {
			sign = SignPositive
			if r == '-' {
				sign = SignNegative
			}
		} else {
			break
		}
	}
	digitsSubstring = strings.TrimLeft(digitsSubstring, "0")
	digitsSubstringLen := len(digitsSubstring)
	result := 0
	if digitsSubstringLen > 10 {
		if sign == SignNegative {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}
	if digitsSubstringLen == 10 {
		if sign == SignNegative && digitsSubstring > MinInt32 {
			return math.MinInt32
		} else if sign == SignPositive && digitsSubstring > MaxInt32 {
			return math.MaxInt32
		}
	}
	for i := 0; i < digitsSubstringLen; i++ {
		digit := digitsMap[rune(digitsSubstring[digitsSubstringLen-i-1])]
		result = result + digit*int(math.Pow(10, float64(i)))
	}
	if sign == SignNegative {
		result = result * -1
	}
	return result
}
