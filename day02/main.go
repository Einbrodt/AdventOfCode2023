package main

import (
	"adventOfCode/day02/game"
	"fmt"
)

func main() {
	filePath := "day02/puzzle.txt"

	cubeBag := game.CubeBag{
		Blue:  14,
		Red:   12,
		Green: 13,
	}

	fmt.Println(game.CubeConundrum(filePath, cubeBag))
	fmt.Println(game.GetFewestCubes(filePath))
}
