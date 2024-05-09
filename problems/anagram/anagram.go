package anagram

import (
	"regexp"
	"strings"
)

func IsAnagram(str1, str2 string) bool {

	regex := regexp.MustCompile(`[^a-z]+`)

	lower1 := strings.ToLower(str1)
	output1 := regex.ReplaceAllString(lower1, "")

	lower2 := strings.ToLower(str2)
	output2 := regex.ReplaceAllString(lower2, "")

	map1 := make(map[byte]int, len(output1))
	map2 := make(map[byte]int, len(output2))

	for i := 0; i < len(output1); i++ {
		map1[output1[i]]++
	}

	for i := 0; i < len(output2); i++ {
		map2[output2[i]]++
	}

	for k, v := range map1 {
		count, ok := map2[k]
		if ok {
			if v != count {
				return false
			}
		}
	}

	return true
}
