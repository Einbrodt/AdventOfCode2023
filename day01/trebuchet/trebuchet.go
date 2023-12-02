package trebuchet

import (
	"adventOfCode/helpers"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var wordToNumber = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

type CalibrationData struct {
	FirstValue string
	LastValue  string
}

func ReadFileByLineAndAddNumbers(filePath string) int {
	var sum int
	var calibrationData []CalibrationData

	wordPattern := `(?:` + strings.Join(getWordPatterns(), `|`) + `)`
	re := regexp.MustCompile(`(?:[0-9]|` + wordPattern + `)`)

	lines, err := helpers.ReadLinesFromFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range lines {
		matches := findAllMatches(re, line)

		firstValue := convertWordToNumber(matches[0])
		lastValue := convertWordToNumber(matches[len(matches)-1])

		calibrationData = append(calibrationData, CalibrationData{FirstValue: firstValue, LastValue: lastValue})
	}

	for _, data := range calibrationData {
		concatValues := data.FirstValue + data.LastValue
		num, _ := strconv.Atoi(concatValues)
		sum += num
	}

	return sum
}

func getWordPatterns() []string {
	var wordPatterns []string
	for word := range wordToNumber {
		wordPatterns = append(wordPatterns, word)
	}
	return wordPatterns
}

func convertWordToNumber(word string) string {
	word = strings.ToLower(word)
	if number, ok := wordToNumber[word]; ok {
		return strconv.Itoa(number)
	}

	return word
}

func findAllMatches(re *regexp.Regexp, input string) []string {
	var matches []string
	for i := 0; i < len(input); i++ {
		substring := input[i:]
		match := re.FindString(substring)
		if match != "" {
			matches = append(matches, match)
		}
	}
	return matches
}
