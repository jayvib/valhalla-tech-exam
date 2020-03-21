package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSumOfTwoItemExistsFrom(t *testing.T) {

	cases := []struct {
		name     string
		input    []int
		sum      int
		isExists bool
	}{
		{
			name:     "Not Exists",
			input:    []int{1, 2, 3, 4},
			sum:      8,
			isExists: false,
		},
		{
			name:     "Exists",
			input:    []int{4, 2, 4, 1},
			sum:      8,
			isExists: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := IsSumOfTwoItemExistsFrom(c.sum, c.input)
			assert.Equal(t, c.isExists, got)
		})
	}
}
