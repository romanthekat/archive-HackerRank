package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Hello, World.")
	scanner.Scan()

	fmt.Print(scanner.Text())
}