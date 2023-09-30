package main

func main() {

}

func EqualFrequency(word string) bool {
	m1 := make(map[rune]int)
	for _, r := range word {
		m1[r] += 1
	}
	if len(m1) == 1 {
		return true
	}
	m2 := make(map[int]int)
	for _, c := range m1 {
		m2[c] += 1
	}
	if len(m2) > 2 {
		return false
	}
	minFreq := len(word)
	maxFreq := 0
	var maxLetters, minLetters []rune
	for r, c := range m1 {
		if c > maxFreq {
			maxFreq = c
			maxLetters = []rune{r}
		} else if c == maxFreq {
			maxLetters = append(maxLetters, r)
		}
		if c < minFreq {
			minFreq = c
			minLetters = []rune{r}
		} else if c == minFreq {
			minLetters = append(minLetters, r)
		}
	}
	return (len(minLetters) == 1 && m1[minLetters[0]] == 1) ||
		(len(maxLetters) == 1 && m1[maxLetters[0]]-m1[minLetters[0]] == 1) ||
		len(maxLetters) == len(word) || len(minLetters) == len(word)
}
