package models

import (
	"strings"
)

var (
	// BadWords is an array of bad words
	BadWords = [15]string{
		"fuck",
		"fucked",
		"fucking",
		"pussy",
		"shit",
		"rape",
		"raped",
		"ass",
		"asshole",
		"basterd",
		"cunt",
		"nipples",
		"hot",
		"sexy",
		"naked",
	}
)

// HasBadLanguage searches for bad language in a string
func HasBadLanguage(s string) (bool, string) {
	lowerCase := strings.ToLower(s)
	for i := 0; i < len(BadWords); i++ {
		badWord := BadWords[i]
		if strings.Contains(lowerCase, badWord) {
			return true, badWord
		}
	}
	return false, ""
}
