package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

const MIN_LETTER_ASCII = 97
const MAX_LETTER_ASCII = 122
const MIN_DECIMAL_ASCII = 48
const MAX_DECIMAL_ASCII = 57

func init() {
	rand.Seed(time.Now().UnixNano())
}

// I expect it to be somewhere around 5mins
// It took me a little bit more with some manual testing -> 15mins
func generate(wantCorrect bool) string {
	str := ""
	strLength := rand.Intn(100) + 3
	procudeDigits := true

	if strLength == 0 {
		return getDigit()
	}

	for strLength > 2 {
		str += "-"
		strLength--
		n := rand.Intn(7) + 1

		i := 0
		for i < n && strLength > 0 {
			if procudeDigits {
				str += getDigit()

			} else {
				str += getLetter()
			}
			strLength--
			i++
		}
		procudeDigits = !procudeDigits

	}

	str = str[1:]

	if !wantCorrect {
		n := rand.Intn(2)
		switch n {
		case 0:
			str += "-"
		case 1:
			i := rand.Intn(len(str))
			runes := []rune(str)
			if unicode.IsDigit(runes[i]) {
				runes[i] = []rune(getLetter())[0]
				str = string(runes)

			} else {
				runes[i] = []rune(getDigit())[0]
				str = string(runes)
			}
		default:
			str = "-" + str
		}

	}

	return str

}

func getLetter() string {
	ascii := rand.Intn(MAX_LETTER_ASCII - MIN_LETTER_ASCII)
	ascii += MIN_LETTER_ASCII

	return fmt.Sprintf("%c", ascii)
}

func getDigit() string {
	ascii := rand.Intn(MAX_DECIMAL_ASCII - MIN_DECIMAL_ASCII)
	ascii += MIN_DECIMAL_ASCII
	return fmt.Sprintf("%c", ascii)
}
