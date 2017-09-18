package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	/*elementsNum :=*/ readElementsCount()
	elementsString := readElementsString()
	elements := parseElements(elementsString)

	fmt.Println(getSum(elements))
}

func getSum(elements []int) int {
	sum := 0

	for _, element := range elements {
		sum += element
	}

	return sum
}

func parseElements(elementsString string) []int {
	splittedElements := strings.Split(elementsString, " ")
	var elements []int

	for _, elementsRaw := range splittedElements {
		elements = append(elements, getElement(elementsRaw))
	}

	return elements
}

func getElement(rawElement string) int {
	element, err := strconv.Atoi(rawElement)
	if err != nil {
		panic(err)
	}

	return element
}

func readElementsCount() int {
	var count int
	fmt.Scanf("%v", &count)

	return count
}

func readElementsString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return strings.TrimSpace(scanner.Text())
}
