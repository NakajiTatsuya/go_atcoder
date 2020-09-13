package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDouble(t *testing.T) {
	asserts := assert.New(t)
	testCases := []struct {
		title  string
		input  int
		answer int
	}{
		{
			title:  "テストケース1:2×2",
			input:  2,
			answer: 4,
		},
		{
			title:  "テストケース2:3×2",
			input:  3,
			answer: 6,
		},
		{
			title:  "テストケース3:9999×2",
			input:  9999,
			answer: 19998,
		},
		{
			title:  "テストケース4:-12×2",
			input:  -12,
			answer: -24,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.title, func(t *testing.T) {
			output := double(testCase.input)
			asserts.Equal(output, testCase.answer)
		})
	}
}
