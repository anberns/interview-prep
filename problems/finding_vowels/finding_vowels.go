package findingvowels

import "strings"

func FindVowels(str string) int {
	lStr := strings.ToLower(str)
	count := 0
	for i := 0; i < len(lStr); i++ {
		c := lStr[i]
		cStr := string(c)
		if cStr == "a" || cStr == "e" || cStr == "i" || cStr == "o" || cStr == "u" {
			count++
		}
	}

	return count
}
