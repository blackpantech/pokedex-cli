package utils

import "strings"

func CapitalizeFirstLetters(phrase string) string {
	phraseSlice := strings.Split(phrase, "-")
	for index, word := range phraseSlice {
		phraseSlice[index] = strings.ToUpper(word[:1]) + word[1:]
	}
	return strings.Join(phraseSlice, " ")
}
