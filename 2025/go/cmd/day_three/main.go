package main

import (
	_ "embed"
	"fmt"
	"maps"
	"os"
	"slices"
	"sort"
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
	var result int = 0
	for _, bank := range strings.Split(inputData, "\n") {
		largestBattery := slices.Max([]rune(bank[0 : len(bank)-1]))
		remainingBatteries := bank[strings.IndexRune(bank, largestBattery)+1:]
		secondLargestBattery := slices.Max([]rune(remainingBatteries))

		joltage, _ := strconv.Atoi(string([]rune{largestBattery, secondLargestBattery}))
		result += joltage
	}

	utilities.PrintResult(result)
}

func partTwo() {
	var result int = 0
	var targetLength int = 12
	for _, bank := range strings.Split(inputData, "\n") {
		uniqueBatteries := make(map[rune]struct{})

		for _, battery := range bank {
			uniqueBatteries[battery] = struct{}{}
		}

		uniqueBatteriesSlice := slices.Collect(maps.Keys(uniqueBatteries))
		sort.Slice(uniqueBatteriesSlice, func(i, j int) bool {
			return uniqueBatteriesSlice[i] > uniqueBatteriesSlice[j]
		})

		largestBatteryCombinationRunes := getLargestBatteryCombination(bank, uniqueBatteriesSlice, targetLength)
		largestBatteryCombination, err := strconv.Atoi(string(largestBatteryCombinationRunes))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		//fmt.Printf("Found largest battery setup: %d\n", largestBatteryCombination)
		result += largestBatteryCombination
	}

	utilities.PrintResult(result)
}

func getLargestBatteryCombination(bank string, uniqueBatteries []rune, maximumBatteriesIncluded int) []rune {
	var largestBatteryCombination []rune
	var largestBatterySeen rune
	var largestBatterySeenIndex int

	if maximumBatteriesIncluded < 1 {
		return nil
	}

	for _, uniqueBattery := range uniqueBatteries {
		largestBatterySeenIndex = strings.IndexRune(bank, uniqueBattery)

		if largestBatterySeenIndex == -1 {
			continue
		}

		if len(bank[largestBatterySeenIndex:]) < maximumBatteriesIncluded {
			continue
		}

		largestBatterySeen = uniqueBattery
		break
	}

	largestBatteryCombination = append(largestBatteryCombination, largestBatterySeen)

	if len(bank) > 0 {
		largestSubBatteryCombination := getLargestBatteryCombination(bank[largestBatterySeenIndex+1:], uniqueBatteries, maximumBatteriesIncluded-1)

		largestBatteryCombination = append(largestBatteryCombination, largestSubBatteryCombination...)
	}

	return largestBatteryCombination
}
