package main

import (
	"adventOfCode/day01/trebuchet"
	"gotest.tools/assert"
	"testing"
)

func TestReadFileByLineAndAddNumbers(t *testing.T) {
	got := trebuchet.ReadFileByLineAndAddNumbers("../day01/puzzle.txt")
	want := 281

	assert.Equal(t, got, want)
}
