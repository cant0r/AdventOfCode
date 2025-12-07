package utilities

import (
	"errors"
	"fmt"
)

func PrintResult(result int) {
	fmt.Printf("%10d\n", result)
}

func Mod(a, b int) int {
	return (a%b + b) % b
}

func Pow(a, b int) (int, error) {
	result := a

	if b < 0 {
		return 0, errors.New("For negative exponent please use math.Pow due to it being float64 comaptible.")
	}

	if b == 0 {
		return 1, nil
	}

	for i := 1; i < b; i++ {
		result *= a
	}

	return result, nil
}

func GetUniqueCharacters(input string) (map[rune]bool, int) {
	uniqueCharacters := make(map[rune]bool)

	for _, char := range input {
		uniqueCharacters[char] = true
	}

	return uniqueCharacters, len(uniqueCharacters)
}
