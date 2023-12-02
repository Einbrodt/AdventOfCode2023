package game

import (
	"adventOfCode/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	GameNumber int
	Blue       int
	Red        int
	Green      int
}

type CubeBag struct {
	Blue  int
	Red   int
	Green int
}

func CubeConundrum(filePath string, cubeBag CubeBag) int {
	var sum int
	games := iterateGames(filePath)
	for _, game := range games {
		if game.Blue <= cubeBag.Blue && game.Red <= cubeBag.Red && game.Green <= cubeBag.Green {
			sum += game.GameNumber
		}
	}
	return sum
}

func GetFewestCubes(filePath string) int {
	var sum int
	games := iterateGames(filePath)
	for _, game := range games {
		sum += game.Red * game.Blue * game.Green
	}
	return sum
}

func iterateGames(filePath string) []Game {
	var games []Game

	lines, err := helpers.ReadLinesFromFile(filePath)
	if err != nil {
		fmt.Println(err)
		return games
	}

	for gameNumber, line := range lines {
		game := parseGame(line)
		game.GameNumber = gameNumber + 1
		games = append(games, game)
	}
	return games
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
