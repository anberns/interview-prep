package main

import (
	"fmt"
)

func main() {
	//r := reverseArr([]int{1, 2, 3, 4, 5})
	arr := []int{1, 2, 3, 4, 5, 6}
	reverseArrInPlace(arr)
	fmt.Println(arr)
}

func reverseArr(arr []int) []int {
	rev := make([]int, len(arr))
	for i, el := range arr {
		rev[len(rev)-1-i] = el
	}

	return rev
}

func reverseArrInPlace(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		temp := arr[len(arr)-1-i]
		arr[len(arr)-1-i] = arr[i]
		arr[i] = temp
	}
}
