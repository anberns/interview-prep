package maxchars

func MaxChar(input string) string {
	var maxChar byte
	maxCharCount := 0
	b := []byte(input)
	charMap := make(map[byte]int, len(input))
	for _, el := range b {
		count, ok := charMap[el]
		if ok {
			charMap[el]++
			if count > maxCharCount {
				maxCharCount = count
				maxChar = el
			}

		} else {
			charMap[el] = 1
		}
	}

	return string(maxChar)
}
