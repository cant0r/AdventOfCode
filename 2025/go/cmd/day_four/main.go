package main

import (
	_ "embed"
	"strings"

	"github.com/cant0r/AdventOfCode/2025/go/internal/utilities"
)

//go:embed input.txt
var inputData string

type Coordinate struct {
	row int
	col int
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	var result int = 0

	paperMap := strings.Split(inputData, "\n")

	for row, paperMapRow := range paperMap {
		for col, _ := range paperMapRow {
			adjecentPaperRolls := 0

			if paperMap[row][col] == '@' {
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						newRow := row + i
						newCol := col + j

						if newRow < 0 || newCol < 0 || newRow >= len(paperMap) || newCol >= len(paperMapRow) || (newRow == row && newCol == col) {
							continue
						}

						if paperMap[newRow][newCol] == '@' {
							adjecentPaperRolls++
						}
					}
				}

				if adjecentPaperRolls < 4 {
					result++
				}
			}
		}
	}

	utilities.PrintResult(result)
}

func partTwo() {
	var result int = 0

	paperMapRows := strings.Split(inputData, "\n")
	paperMap := make([][]rune, len(paperMapRows))

	for row, paperMapRow := range paperMapRows {
		paperMap[row] = []rune(paperMapRow)
	}

	for {
		var roundResult int = 0
		var approachablePaperRools []Coordinate

		for row, paperMapRow := range paperMap {
			for col, _ := range paperMapRow {
				adjecentPaperRolls := 0

				if paperMap[row][col] == '@' {
					for i := -1; i <= 1; i++ {
						for j := -1; j <= 1; j++ {
							newRow := row + i
							newCol := col + j

							if newRow < 0 || newCol < 0 || newRow >= len(paperMap) || newCol >= len(paperMapRow) || (newRow == row && newCol == col) {
								continue
							}

							if paperMap[newRow][newCol] == '@' {
								adjecentPaperRolls++
							}
						}
					}

					if adjecentPaperRolls < 4 {
						roundResult++
						approachablePaperRools = append(approachablePaperRools, Coordinate{row, col})
					}
				}
			}
		}

		for _, approachablePaperRoll := range approachablePaperRools {
			paperMap[approachablePaperRoll.row][approachablePaperRoll.col] = 'x'
		}

		//fmt.Printf("Round result: %d\n", roundResult)
		if roundResult == 0 {
			break
		}

		result += roundResult
	}

	utilities.PrintResult(result)
}
