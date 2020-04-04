package main

import (
	"sort"
)

func permuteUnique(nums []int) (permutations [][]int) {
	sort.Ints(nums)
	if len(nums) == 0 {
		return
	}
	if len(nums) == 1 {
		permutations = append(permutations, nums)
		return
	}

	uniques := getUniques(nums)
	for _, unique := range uniques {
		a := make([]int, len(nums))
		copy(a, nums)

		remainder := append(a[:unique.i], a[unique.i+1:]...)
		for _, permutation := range permuteUnique(remainder) {
			permutations = append(permutations, append([]int{unique.num}, permutation...))
		}
	}

	return
}

// Unique represents the index and value of a unique number in a slice
type Unique struct {
	i, num int
}

func getUniques(nums []int) (uniques []Unique) {
	prevNum := nums[0] - 1
	for i, num := range nums {
		if num > prevNum {
			uniques = append(uniques, Unique{i: i, num: num})
		}
		prevNum = num
	}
	return
}

func main() {
	permuteUnique([]int{1, 2, 3})
}
