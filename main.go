package main

import (
	"math"
	"strings"
	"unicode"
)

// Estimate: 5mins
// Real: <10mins
func testValidity(str string) bool {
	isDigitsSlot := true
	l := len(str)

	// this is safe becase we assume the string received has length larger than 0
	firstChar := []rune(str)[0]
	if !unicode.IsDigit(firstChar) {
		return false
	}

	for i, char := range str {
		isDigit := unicode.IsDigit(char)
		isLetter := unicode.IsLetter(char)
		if i == l-1 && !isDigit && !isLetter {
			return false
		}

		isDash := string(char) == "-"
		if isDash {
			isDigitsSlot = !isDigitsSlot
			continue
		}

		if isDigit && !isDigitsSlot || !isDigit && isDigitsSlot {
			return false
		}
	}

	return true
}

// Estimate: <5mins
// Real: <5mins
func averageNumber(str string) int {
	n := 0
	sum := 0

	for _, char := range str {
		if unicode.IsDigit(char) {
			sum += int(char - '0')
			n++
		}
	}

	if n == 0 {
		return 0
	}
	return sum / n
}

// Estimate: <5 mins
// Real: <5mins
func wholeStory(str string) string {
	res := ""
	isLetters := false

	for _, char := range str {

		isDash := string(char) == "-"
		if isDash && isLetters {
			res += " "
			isLetters = false
			continue
		}
		if unicode.IsLetter(char) {
			isLetters = true
			res += string(char)
		}
	}

	return res
}

type Stats struct {
	shortestWord      string
	longestWord       string
	avgWordLength     int
	avgWordLengthList []string
}

// Estimate: 5-10mins
// Real: 5-10mins
func storyStats(str string) *Stats {
	stats := &Stats{}
	maxLength := 0
	minLength := math.MaxInt
	sum := 0

	wholeStory := wholeStory(str)
	words := strings.Split(wholeStory, " ")
	for _, w := range words {
		l := len(w)
		sum += l
		if l < minLength {
			minLength = l
			stats.shortestWord = w
		}
		if l > maxLength {
			maxLength = l
			stats.longestWord = w
		}
	}
	if l := len(words); l > 0 {
		stats.avgWordLength = sum / len(words)

		for _, w := range words {
			if len(w) == stats.avgWordLength {
				stats.avgWordLengthList = append(stats.avgWordLengthList, w)
			}
		}
	}

	return stats

}
