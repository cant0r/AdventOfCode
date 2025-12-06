package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"

	"github.com/cant0r/AdventOfCode/2025/go/internal/utilities"
)

//go:embed input.txt
var inputData string

type Range struct {
	Start int
	End   int
}

func (rng *Range) Overlaps(other *Range) bool {
	if rng.Start <= other.Start && other.Start <= rng.End {
		return true
	}

	if rng.Start <= other.End && other.End <= rng.End {
		return true
	}

	if other.Start <= rng.Start && rng.Start <= other.End {
		return true
	}

	if other.Start <= rng.End && rng.End <= other.End {
		return true
	}

	return false
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	var result int = 0
	var ranges []Range

	for _, dataRow := range strings.Split(inputData, "\n") {
		if strings.ContainsRune(dataRow, '-') {
			rangeData := strings.Split(dataRow, "-")
			rangeStart, _ := strconv.Atoi(rangeData[0])
			rangeEnd, _ := strconv.Atoi(rangeData[1])
			ranges = append(ranges, Range{Start: rangeStart, End: rangeEnd})
		} else if strings.TrimSpace(dataRow) == "" {
			continue
		} else {
			id, _ := strconv.Atoi(dataRow)

			for _, idRange := range ranges {
				if idRange.Start <= id && id <= idRange.End {
					result++
					break
				}
			}
		}
	}

	utilities.PrintResult(result)
}

func partTwo() {
	var result int = 0
	var ranges []Range

	for _, dataRow := range strings.Split(inputData, "\n") {
		if strings.ContainsRune(dataRow, '-') {
			rangeData := strings.Split(dataRow, "-")
			rangeStart, _ := strconv.Atoi(rangeData[0])
			rangeEnd, _ := strconv.Atoi(rangeData[1])
			ranges = append(ranges, Range{Start: rangeStart, End: rangeEnd})

		} else {
			break
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	for i := 0; i < len(ranges); i++ {
		for j := i + 1; j < len(ranges); j++ {
			if ranges[i].Overlaps(&ranges[j]) {
				newRangeEnd := 0
				if ranges[i].End > ranges[j].End {
					newRangeEnd = ranges[i].End
				} else {
					newRangeEnd = ranges[j].End
				}
				ranges[i].End = ranges[j].Start
				ranges[j].Start = ranges[j].Start + 1
				ranges[j].End = newRangeEnd
			}
		}
	}

	for _, rng := range ranges {
		result += (rng.End - rng.Start + 1)
	}

	utilities.PrintResult(result)
}
