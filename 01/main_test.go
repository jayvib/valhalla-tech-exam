package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseSentence(t *testing.T) {

	testcases := []struct {
		input string
		want  string
	}{
		{"Hello World", "World Hello"},
		{"Hello    World", "World Hello"},
	}

	for _, tc := range testcases {
		got := reverseSentence(tc.input)
		assert.Equal(t, tc.want, got)
	}

}
