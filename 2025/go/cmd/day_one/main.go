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
	dialPosition := 50
	minimumDial := 0
	maximumDial := 99
	minimumDialHitCount := 0

	for _, line := range strings.Split(inputData, "\n") {
		direction := string(line[0])
		rotation, _ := strconv.Atoi(string(line[1:]))

		if direction == "L" {
			dialPosition = utilities.Mod(dialPosition-rotation, maximumDial+1)
		} else {
			dialPosition = utilities.Mod(dialPosition+rotation, maximumDial+1)
		}

		//fmt.Printf("Now at %d\n", dialPosition)

		if dialPosition == minimumDial {
			minimumDialHitCount += 1
		}
	}

	utilities.PrintResult(minimumDialHitCount)
}

func partTwo() {
	dialPosition := 50
	minimumDial := 0
	maximumDial := 100
	minimumDialHitCount := 0

	for _, line := range strings.Split(inputData, "\n") {
		direction := string(line[0])
		rotation, _ := strconv.Atoi(string(line[1:]))

		wentPassCount := rotation / maximumDial

		if direction == "L" {
			if nextDialPosition := dialPosition - utilities.Mod(rotation, maximumDial); dialPosition != 0 && nextDialPosition < minimumDial {
				wentPassCount += 1
			}
			dialPosition = utilities.Mod(dialPosition-rotation, maximumDial)
		} else {
			if nextDialPosition := dialPosition + utilities.Mod(rotation, maximumDial); nextDialPosition > maximumDial {
				wentPassCount += 1
			}
			dialPosition = utilities.Mod(dialPosition+rotation, maximumDial)
		}

		minimumDialHitCount += wentPassCount

		if dialPosition == minimumDial {
			minimumDialHitCount += 1
		}
	}

	utilities.PrintResult(minimumDialHitCount)
}
