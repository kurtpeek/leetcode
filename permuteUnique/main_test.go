package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermuteUnique(t *testing.T) {
	var testCases = []struct {
		input  []int
		output [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1, 1, 2}, [][]int{[]int{1, 1, 2}, []int{1, 2, 1}, []int{2, 1, 1}}},
		{[]int{1, 2, 3}, [][]int{[]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}, []int{2, 3, 1}, []int{3, 1, 2}, []int{3, 2, 1}}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			assert.Exactly(t, tc.output, permuteUnique(tc.input))
		})
	}
}

func TestGetUniques(t *testing.T) {
	var testCases = []struct {
		nums    []int
		uniques []Unique
	}{
		{[]int{1, 1, 2}, []Unique{Unique{0, 1}, Unique{2, 2}}},
		{[]int{1, 2, 3}, []Unique{Unique{0, 1}, Unique{1, 2}, Unique{2, 3}}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			assert.Exactly(t, tc.uniques, getUniques(tc.nums))
		})
	}
}
