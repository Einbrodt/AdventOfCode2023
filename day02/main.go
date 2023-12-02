package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	GameNumber int
	Blue       int
	Red        int
	Green      int
}

func CubeConundrum(filePath string) int {
	var sum int
	var games []Game
	cubeBag := Game{
		Blue:  14,
		Red:   12,
		Green: 13,
	}

	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 0
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for gameNumber := 1; fileScanner.Scan(); gameNumber++ {
		game := parseGame(fileScanner.Text())
		game.GameNumber = gameNumber
		games = append(games, game)
	}

	for _, game := range games {
		if game.Blue <= cubeBag.Blue && game.Red <= cubeBag.Red && game.Green <= cubeBag.Green {
			sum += game.GameNumber
			fmt.Printf("Game %d has been added.\t", game.GameNumber)
		}
		fmt.Printf("Game %d: Blue=%d, Red=%d, Green=%d\n", game.GameNumber, game.Blue, game.Red, game.Green)
	}

	fmt.Println(sum)
	return sum
}

func parseGame(input string) Game {
	game := Game{
		Blue:  0,
		Red:   0,
		Green: 0,
	}

	colonIndex := strings.Index(input, ":")
	if colonIndex == -1 {
		return game
	}

	drawsString := input[colonIndex+1:]
	draws := strings.Split(drawsString, ";")
	maxCounts := make(map[string]int)

	for _, draw := range draws {
		colorCounts := strings.Split(draw, ",")
		roundCounts := make(map[string]int)

		for _, colorCount := range colorCounts {
			parts := strings.Fields(colorCount)
			if len(parts) == 2 {
				count, err := strconv.Atoi(parts[0])
				if err != nil {
					fmt.Println("Error converting count to integer:", err)
					continue
				}

				roundCounts[parts[1]] += count
			}
		}

		for color, count := range roundCounts {
			maxCounts[color] = max(maxCounts[color], count)
		}
	}

	game.Blue = maxCounts["blue"]
	game.Red = maxCounts["red"]
	game.Green = maxCounts["green"]

	return game
}
