package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type searchFunc func([]int, int) int

func TestSearchRangeFuncs(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}

	var testCases = []struct {
		name             string
		nums             []int
		searchFunc       searchFunc
		target, expected int
	}{
		{"left1", nums, searchRangeLeft, 7, 1},
		{"left2", nums, searchRangeLeft, 8, 3},
		{"leftNoMatch1", nums, searchRangeLeft, 4, -1},
		{"leftNoMatch2", nums, searchRangeLeft, 9, -1},
		{"leftNoMatch3", nums, searchRangeLeft, 11, -1},
		{"leftmost", nums, searchRangeLeft, 5, 0},
		{"right1", nums, searchRangeRight, 7, 2},
		{"right2", nums, searchRangeRight, 8, 4},
		{"rightNoMatch1", nums, searchRangeRight, 4, -1},
		{"rightNoMatch2", nums, searchRangeRight, 9, -1},
		{"rightNoMatch3", nums, searchRangeRight, 11, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Exactly(t, tc.expected, tc.searchFunc(nums, tc.target))
		})
	}
}

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}

	var testCases = []struct {
		nums    []int
		target  int
		indices []int
	}{
		{nums, 8, []int{3, 4}},
		{nums, 7, []int{1, 2}},
		{nums, 4, []int{-1, -1}},
		{nums, 10, []int{5, 5}},
		{[]int{1, 1}, 1, []int{0, 1}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			assert.Exactly(t, tc.indices, searchRange(tc.nums, tc.target))
		})
	}
}
