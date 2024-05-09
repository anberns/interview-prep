package palindrome

import ra "example.com/reverse_array"

func IsPalindrome(input string) bool {
	rev := ra.ReverseArrInPlaceMostly(input)

	return input == rev
}

func IsPalindromeInPlace(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}

	return true
}
