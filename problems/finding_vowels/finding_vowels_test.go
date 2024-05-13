package findingvowels_test

import (
	"testing"

	findingvowels "example.com/finding_vowels"
	"github.com/stretchr/testify/require"
)

func TestFindVowels(t *testing.T) {
	str := "A somewhat long string has been strung."

	result := findingvowels.FindVowels(str)
	require.Equal(t, 10, result)

	str = "trgh pw"

	result = findingvowels.FindVowels(str)
	require.Equal(t, 0, result)
}
