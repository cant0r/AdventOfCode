package utilities

import "fmt"

func PrintResult(result int) {
	fmt.Printf("%10d\n", result)
}

func Mod(a, b int) int {
	return (a%b + b) % b
}
