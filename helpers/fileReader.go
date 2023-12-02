package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLinesFromFile(filePath string) ([]string, error) {
	var lines []string

	file, err := os.Open(filePath)
	if err != nil {
		return lines, fmt.Errorf("error opening file: %w", err)
	}
	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return lines, fmt.Errorf("error scanning file: %w", err)
	}

	return lines, nil
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
	}
}
