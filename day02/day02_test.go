package day02

import (
	"gotest.tools/assert"
	"testing"
)

func TestForExamples(t *testing.T) {
	got := CubeConundrum("../day02/examples.txt")
	want := 8

	assert.Equal(t, got, want)
}
