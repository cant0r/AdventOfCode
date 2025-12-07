package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/cant0r/AdventOfCode/2025/go/internal/utilities"
)

//go:embed input.txt
var inputData string

func main() {
	partOne()
	partTwo()
}

func printTeleportMap(teleportMap [][]rune) {
	for _, row := range teleportMap {
		fmt.Printf("%q\n", row)
	}
}

func partOne() {
	var result int = 0

	teleportMapRows := strings.Split(inputData, "\n")
	teleportMap := make([][]rune, len(teleportMapRows))

	for i, row := range teleportMapRows {
		teleportMap[i] = []rune(row)
	}

	for i, row := range teleportMap {
		for j, col := range row {
			if col == 'S' {
				teleportMap[i+1][j] = '|'
				continue
			}

			if col == '^' && teleportMap[i-1][j] == '|' {
				split := false
				if teleportMap[i][j-1] != '|' {
					teleportMap[i][j-1] = '|'
					split = true
				}
				if teleportMap[i][j+1] != '|' {
					teleportMap[i][j+1] = '|'
					split = true
				}

				if split {
					result++
				}

			} else {
				if i > 0 && teleportMap[i-1][j] == '|' {
					teleportMap[i][j] = '|'
				}
			}
		}
	}

	//printTeleportMap(teleportMap)

	utilities.PrintResult(result)
}

var beamPathCache map[string]int = make(map[string]int)

func findNumberOfPaths(teleportMap [][]rune, rowPosition, beamColPosiition int) int {
	if inCache := beamPathCache[fmt.Sprintf("%d-%d", rowPosition, beamColPosiition)]; inCache != 0 {
		return inCache
	}

	cacheEntry := fmt.Sprintf("%d-%d", rowPosition, beamColPosiition)

	if rowPosition == len(teleportMap)-1 {
		beamPathCache[cacheEntry] = 1
		return beamPathCache[cacheEntry]
	}

	if teleportMap[rowPosition+1][beamColPosiition] != '^' {
		beamPathCache[cacheEntry] = findNumberOfPaths(teleportMap, rowPosition+1, beamColPosiition)
		return beamPathCache[cacheEntry]
	}

	if teleportMap[rowPosition+1][beamColPosiition] == '^' {
		beamPathCache[cacheEntry] = findNumberOfPaths(teleportMap, rowPosition+1, beamColPosiition-1) + findNumberOfPaths(teleportMap, rowPosition+1, beamColPosiition+1)
		return beamPathCache[cacheEntry]
	}
	return 0
}

func partTwo() {
	var result int = 0

	teleportMapRows := strings.Split(inputData, "\n")
	teleportMap := make([][]rune, len(teleportMapRows))

	for i, row := range teleportMapRows {
		teleportMap[i] = []rune(row)
	}

	firstBeamCol := 0

	for i, col := range teleportMap[0] {
		if col == 'S' {
			firstBeamCol = i
			break
		}
	}

	result += findNumberOfPaths(teleportMap, 1, firstBeamCol)

	utilities.PrintResult(result)
}
