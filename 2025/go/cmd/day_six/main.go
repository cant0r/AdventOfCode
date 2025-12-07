package main

import (
	_ "embed"
	"strconv"
	"strings"
	"sync"

	"github.com/cant0r/AdventOfCode/2025/go/internal/utilities"
)

//go:embed input.txt
var inputData string

func main() {
	partOne()
	partTwo()
}

func loadExercies(inputData string) (opers [][]int, ops []string) {
	inputDataRows := strings.Split(inputData, "\n")
	operators := strings.Fields(inputDataRows[len(inputDataRows)-1])
	inputDataRows = inputDataRows[:len(inputDataRows)-1]
	numberOfExercies := len(strings.Fields(inputDataRows[0]))
	operands := make([][]int, numberOfExercies)

	for i := 0; i < numberOfExercies; i++ {
		operands[i] = make([]int, len(inputDataRows))
	}

	for i, subOperands := range inputDataRows {
		for j, exerciseOperand := range strings.Fields(subOperands) {
			operands[j][i], _ = strconv.Atoi(exerciseOperand)
		}
	}

	return operands, operators
}

func partOne() {
	var result int = 0
	var wg sync.WaitGroup

	operands, operators := loadExercies(inputData)
	exerciseChannel := make(chan int, len(operators))

	for i, operator := range operators {
		switch operator {
		case "+":
			wg.Go(func() {
				subResult := 0
				for _, operand := range operands[i] {
					subResult += operand
				}
				exerciseChannel <- subResult
			})
		case "*":
			wg.Go(func() {
				subResult := 1
				for _, operand := range operands[i] {
					subResult *= operand
				}
				exerciseChannel <- subResult
			})
		}
	}

	wg.Wait()

	for i := 0; i < len(operators); i++ {
		result += <-exerciseChannel
	}

	utilities.PrintResult(result)
}

func loadOctopusExercies(inputData string) (opers map[int][]int, ops []string) {
	inputDataRows := strings.Split(inputData, "\n")
	operators := strings.Fields(inputDataRows[len(inputDataRows)-1])
	inputDataRows = inputDataRows[:len(inputDataRows)-1]
	numberOfExercies := len(strings.Fields(inputDataRows[0]))
	operands := make(map[int][]int, numberOfExercies)

	exerciseCounter := 0

	for i := 0; i < len(inputDataRows[0]); i++ {
		var operand strings.Builder
		for j := 0; j < len(inputDataRows); j++ {

			if inputDataRows[j][i] != ' ' {
				operand.WriteByte(inputDataRows[j][i])
			}
		}

		if len(strings.TrimSpace(operand.String())) == 0 {
			exerciseCounter++
			continue
		}

		newOperand, _ := strconv.Atoi(operand.String())
		if exisitingOperands := operands[exerciseCounter]; exisitingOperands == nil {
			operands[exerciseCounter] = []int{newOperand}
		} else {
			operands[exerciseCounter] = append(operands[exerciseCounter], newOperand)
		}
	}

	return operands, operators
}

func partTwo() {
	var result int = 0
	var wg sync.WaitGroup

	operands, operators := loadOctopusExercies(inputData)
	exerciseChannel := make(chan int, len(operators))

	for i, operator := range operators {
		//fmt.Printf("Got octopus numbers %v\n", operands)
		switch operator {
		case "+":
			wg.Go(func() {
				subResult := 0
				for _, operand := range operands[i] {
					subResult += operand
				}
				exerciseChannel <- subResult
			})
		case "*":
			wg.Go(func() {
				subResult := 1
				for _, operand := range operands[i] {
					subResult *= operand
				}
				exerciseChannel <- subResult
			})
		}
	}

	wg.Wait()

	for i := 0; i < len(operators); i++ {
		result += <-exerciseChannel
	}

	utilities.PrintResult(result)
}
