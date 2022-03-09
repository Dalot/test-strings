package main

import (
	"unicode"
)

// I assumed it was going to take somewhat around 5mins. And it did the first version. Then I reiterated on it so the total must have been between 5-10mins
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

// I assumed this one would be faster than the first one so, <5mins
// Yes, this one took probably 1min
func averageNumber(str string) int {
	n := 0
	sum := 0

	for _, char := range str {
		if unicode.IsDigit(char) {
			sum += int(char)
			n++
		}
	}

	if n == 0 {
		return 0
	}
	return sum / n
}

// I assume this one will take <5 mins
// It tooks similar time than the averageNumber
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
