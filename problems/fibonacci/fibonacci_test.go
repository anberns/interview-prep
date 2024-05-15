package fibonacci_test

import (
	"testing"

	"example.com/fibonacci"
	"github.com/stretchr/testify/require"
)

func TestFibIter(t *testing.T) {

	num := 10
	result := fibonacci.FibIter(num)
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	require.Equal(t, expected, result)

	num = 0
	result = fibonacci.FibIter(num)
	expected = []int{}
	require.Equal(t, expected, result)

	num = 1
	result = fibonacci.FibIter(num)
	expected = []int{0}
	require.Equal(t, expected, result)

	num = 2
	result = fibonacci.FibIter(num)
	expected = []int{0, 1}
	require.Equal(t, expected, result)
}

func TestFibRec(t *testing.T) {

	n := 9
	result := fibonacci.FibRec(n)
	expected := 34
	require.Equal(t, expected, result)

	n = 0
	result = fibonacci.FibRec(n)
	expected = 0
	require.Equal(t, expected, result)

	n = 1
	result = fibonacci.FibRec(n)
	expected = 1
	require.Equal(t, expected, result)

	n = 2
	result = fibonacci.FibRec(n)
	expected = 1
	require.Equal(t, expected, result)
}
