package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

const LOCK_LENGTH = 5

func main() {
	inputLines := readInputLines()
	movesCount := calculateMovesNumber(inputLines)

	fmt.Println(movesCount)
}

func calculateMovesNumber(inputLines []string) int {
	totalMovements := 0
	currentConfig, desiredConfig := parseConfigurations(inputLines)

	for i := 0; i < LOCK_LENGTH; i++ {
		currentValue := currentConfig[i]
		desiredValue := desiredConfig[i]

		totalMovements += movesTo(currentValue, desiredValue)
	}

	return totalMovements
}

func movesTo(positionFrom int, positionTo int) int {
	if positionFrom == positionTo {
		return 0
	}

	positiveMoves := positiveMoves(positionFrom, positionTo)
	negativeMoves := negativeMoves(positionFrom, positionTo)

	if positiveMoves > negativeMoves {
		return negativeMoves
	} else {
		return positiveMoves
	}
}

func negativeMoves(positionFrom int, positionTo int) int {
	moves := 0

	for {
		if positionFrom == positionTo {
			break
		}

		positionFrom -= 1
		if positionFrom == -1 {
			positionFrom = 9
		}

		moves++
	}

	return moves
}

func positiveMoves(positionFrom int, positionTo int) int {
	moves := 0

	for {
		if positionFrom == positionTo {
			break
		}

		positionFrom += 1
		if positionFrom == 10 {
			positionFrom = 0
		}

		moves++
	}

	return moves
}

func parseConfigurations(inputLines []string) ([]int, []int) {
	currentConfig := parseConfiguration(inputLines[0])
	desiredConfig := parseConfiguration(inputLines[1])

	return currentConfig, desiredConfig
}

func parseConfiguration(inputLine string) []int {
	config := make([]int, LOCK_LENGTH)

	configStr := strings.Split(inputLine, " ")
	for index, numberAsStr := range configStr {
		number, err := strconv.Atoi(numberAsStr)
		if err != nil {
			panic(err)
		}

		config[index] = number
	}

	return config
}

func readInputLines() []string {
	scanner := bufio.NewScanner(os.Stdin)

	inputLines := make([]string, 2)
	for i := 0; i < 2; i++ {
		scanner.Scan()

		err := scanner.Err()
		if err != nil {
			panic(err)
		}

		inputLines[i] = scanner.Text()
	}

	return inputLines
}