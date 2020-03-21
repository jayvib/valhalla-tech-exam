package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHighestAndLowestItem(t *testing.T) {

	cases := []struct {
		input           []int
		highest, lowest int
	}{
		{
			input:   []int{2, 1, 6, 7, 3},
			highest: 7,
			lowest:  1,
		},
		{
			input:   []int{2, 6, 8, 3},
			highest: 8,
			lowest:  2,
		},
	}

	for _, c := range cases {
		lowest, highest := GetHighestAndLowestItem(c.input)
		assert.Equal(t, c.lowest, lowest)
		assert.Equal(t, c.highest, highest)
	}
}
