package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFirstRecurringCharacter(t *testing.T) {

	cases := []struct {
		name  string
		input string
		want  rune
		err   error
	}{
		{name: "First Occurence Duplicate Char returns A", input: "ABCA", want: 'A'},
		{name: "First Occurence Duplicate Char returns B", input: "BACB", want: 'B'},
		{name: "First Occurence Duplicate Not Found", input: "CBAD", want: ' ', err: ErrNotFound},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := GetFirstRecurringCharacter(c.input)
			assert.Equal(t, c.err, err)
			assert.Equal(t, c.want, got)
		})
	}

}
