package main

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/cant0r/AdventOfCode/2025/go/internal/utilities"
)

//go:embed input.txt
var inputData string

func main() {
	partOne()
	partTwo()
}

func partOne() {
	result := 0
	for _, idRange := range strings.Split(inputData, ",") {
		idRangeParts := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(idRangeParts[0])
		end, _ := strconv.Atoi(idRangeParts[1])

		for i := start; i <= end; i++ {
			iString := strconv.Itoa(i)
			prefix := iString[0 : len(iString)/2]
			suffix := strings.TrimPrefix(iString, prefix)

			if prefix == suffix {
				//fmt.Printf("Found an invalid id = %d\n", i)
				result += i
			}
		}
	}

	utilities.PrintResult(result)
}

func partTwo() {

	result := 0
	for _, idRange := range strings.Split(inputData, ",") {
		idRangeParts := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(idRangeParts[0])
		end, _ := strconv.Atoi(idRangeParts[1])

		for i := start; i <= end; i++ {
			iString := strconv.Itoa(i)

			for endPointer := 1; endPointer <= len(iString)/2; endPointer++ {
				prefix := iString[0:endPointer]
				prefixOccurences := strings.Count(iString, prefix)

				if prefixOccurences >= 2 && len(iString) == prefixOccurences*len(prefix) {
					//fmt.Printf("Found an invalid id = %d with prefix=%s with prefixOccurences=%d\n", i, prefix, prefixOccurences)
					result += i
					break
				}
			}
		}
	}

	utilities.PrintResult(result)
}
