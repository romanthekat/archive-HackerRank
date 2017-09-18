package main

import (
	"fmt"
	"os"
)

func main() {
	sum := 0
	elementsCount := readElementsCount()

	var element int
	for i := 0; i < elementsCount; i++ {
		_, err := fmt.Fscan(os.Stdin, &element)

		if err != nil {
			break
		}

		sum += element
	}

	fmt.Println(sum)
}

func readElementsCount() int {
	var count int
	fmt.Scanf("%v", &count)

	return count
}
