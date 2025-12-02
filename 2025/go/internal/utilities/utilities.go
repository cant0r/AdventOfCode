package utilities

import "fmt"

func PrintResult(result int) {
	fmt.Printf("%10d\n", result)
}

func Mod(a, b int) int {
	return (a%b + b) % b
}

func GetUniqueCharacters(input string) (map[rune]bool, int) {
	uniqueCharacters := make(map[rune]bool)

	for _, char := range input {
		uniqueCharacters[char] = true
	}

	return uniqueCharacters, len(uniqueCharacters)
}
