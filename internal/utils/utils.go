package utils

import (
	"strings"
)

func CapitalizeFirstLetters(phrase string) string {
	if len(phrase) == 0 {
		return ""
	}
	phraseSlice := strings.Split(phrase, "-")
	for index, word := range phraseSlice {
		phraseSlice[index] = strings.ToUpper(word[:1]) + word[1:]
	}
	return strings.Join(phraseSlice, " ")
}
