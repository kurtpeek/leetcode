package main

import "sort"

func searchRange(nums []int, target int) []int {
	return []int{searchRangeLeft(nums, target), searchRangeRight(nums, target)}
}

func searchRangeLeft(nums []int, target int) int {
	if i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	}); i < len(nums) && nums[i] == target {
		return i
	}
	return -1
}

func searchRangeRight(nums []int, target int) int {
	if i := sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	}); i > 0 && nums[i-1] == target {
		return i - 1
	}
	return -1
}
