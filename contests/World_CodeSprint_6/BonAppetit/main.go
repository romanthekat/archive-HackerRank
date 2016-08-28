package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Input struct {
	itemsCount    int
	allergicIndex int
	items         []int
	charged       int
}

type Output struct {
	overcharge int
}

func main() {
	output := analyzeInput(getInput())

	if output.overcharge == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(output.overcharge)
	}
}

func analyzeInput(input Input) Output {

}

func getInput() Input{
	lines := readInputLines()

	itemsCount, allergicIndex := parseFirstLine(lines[0])
	items := parseSecondLine(lines[1], itemsCount)
	charged := parseThirdLine(lines[2])

	return Input{itemsCount, allergicIndex, items, charged}
}

func readInputLines() []string {
	reader := bufio.NewReader(os.Stdin)

	inputLines := make([]string, 3)
	for i := 0; i < 3; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		inputLines[i] = line
	}

	return inputLines
}

//third line contains number - Brain charged Anna for her share of the bill
//for example '7'
func parseThirdLine(line string) int {
	return int(line)
}

//second line contains spaced numbers - costs of items
//for example '3 10 2 9'
func parseSecondLine(line string, itemsCount int) []int {
	inputItems := strings.Split(line, " ")
	if len(inputItems) != itemsCount {
		panic(fmt.Sprintf("items count doesn't equal to input items count, items:%s, input items count:%s", inputItems, itemsCount))
	}

	results := make([]int, itemsCount)
	for index, inputItem := range inputItems {
		results[index] = int(inputItem)
	}

	return results
}


//first line contains two spaced numbers - itemsCount and allergicIndex,
//for example '4 1'
func parseFirstLine(line string) (int, int) {
	results := strings.Split(line, " ")
	if len(results) != 2 {
		panic(fmt.Sprintf("first input string doesn't contain 2 numbers! input string:%s", line))
	}

	return int(results[0]), int(results[1])
}