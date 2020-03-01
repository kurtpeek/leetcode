package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	var testCases = []struct {
		input, output []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{2, 1, 3}, []int{2, 3, 1}},
		{[]int{2, 3, 1}, []int{3, 1, 2}},
		{[]int{3, 1, 2}, []int{3, 2, 1}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			nextPermutation(tc.input)
			assert.Exactly(t, tc.output, tc.input)
		})
	}
}
