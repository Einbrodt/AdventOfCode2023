package main

import (
	"adventOfCode/day01/trebuchet"
	"fmt"
)

func main() {
	filePath := "day01/puzzle.txt"

	fmt.Println(trebuchet.ReadFileByLineAndAddNumbers(filePath))
}
