package main

import (
	"adventOfCode/day02/game"
	"gotest.tools/assert"
	"testing"
)

func TestCubeConundrum(t *testing.T) {
	cubeBag := game.CubeBag{
		Blue:  14,
		Red:   12,
		Green: 13,
	}
	got := game.CubeConundrum("../day02/examples.txt", cubeBag)
	want := 8

	assert.Equal(t, got, want)
}

func TestGetFewestCubes(t *testing.T) {
	got := game.GetFewestCubes("../day02/examples.txt")
	want := 2286

	assert.Equal(t, got, want)
}
