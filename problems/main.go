package main

import (
	"fmt"

	binarysearch "example.com/binary_search"
)

func main() {
	//reverse_array.ReverseArray()
	// fmt.Println(palindrome.IsPalindromeInPlace("racecar"))
	// fmt.Println(palindrome.IsPalindromeInPlace("aiddia"))
	// fmt.Println(palindrome.IsPalindromeInPlace("aid did"))
	// fmt.Println(palindrome.IsPalindromeInPlace("a"))
	//fmt.Println(maxchars.MaxChar("00011112334555"))
	//fizzbuzz.Fizzbuzz(31)
	//fmt.Println(anagram.IsAnagram("Hi There", "Bye there"))
	fmt.Println(binarysearch.BinarySearch([]int{1, 2, 4, 5, 7, 8, 10, 12}, 1))
	fmt.Println(binarysearch.BinarySearch([]int{1, 2, 4, 5, 7, 8, 10, 12}, 12))
	fmt.Println(binarysearch.BinarySearch([]int{1, 2, 4, 5, 7, 8, 10, 12}, 14))
	fmt.Println(binarysearch.BinarySearch([]int{1, 2, 4, 5, 7, 8, 10, 12}, -1))
	fmt.Println(binarysearch.BinarySearch([]int{1, 2, 4, 5, 7, 8, 10, 12}, 5))
	fmt.Println(binarysearch.BinarySearch([]int{1, 2, 4, 5, 7, 8, 10, 12}, 7))
}
