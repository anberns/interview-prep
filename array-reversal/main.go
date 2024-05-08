package main

import (
	"fmt"
)

func main() {
	//r := reverseArr([]int{1, 2, 3, 4, 5})
	//arr := []int{1, 2, 3, 4, 5, 6}
	//reverseIntArrInPlace(arr)
	str := "Hello World!"
	rev := reverseArrInPlaceMostly(str)
	fmt.Println(rev)
}

func reverseIntArr(arr []int) []int {
	rev := make([]int, len(arr))
	for i, el := range arr {
		rev[len(rev)-1-i] = el
	}

	return rev
}

func reverseIntArrInPlace(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		temp := arr[len(arr)-1-i]
		arr[len(arr)-1-i] = arr[i]
		arr[i] = temp
	}
}

func reverseArrInPlaceMostly(arr string) string {
	bArr := []byte(arr)
	for i := 0; i < len(bArr)/2; i++ {
		temp := arr[len(bArr)-1-i]
		bArr[len(bArr)-1-i] = bArr[i]
		bArr[i] = temp
	}

	return string(bArr)
}
