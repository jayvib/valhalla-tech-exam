package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	input := []int{3, 2, 5, 8, 1, 2}
	want := []int{1, 2, 2, 3, 5, 8}
	bubbleSortInt(input)
	assert.Equal(t, want, input)
}
