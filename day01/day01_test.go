package day01

import (
	"gotest.tools/assert"
	"testing"
)

func TestReadFileByLineAndAddNumbers(t *testing.T) {
	got := ReadFileByLineAndAddNumbers("../day01/puzzle.txt")
	want := 281

	assert.Equal(t, got, want)
}
