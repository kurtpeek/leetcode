package main

import (
	"sort"
)

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			sort.Ints(nums[i:])

			j := sort.Search(len(nums), func(j int) bool {
				return j >= i && nums[j] > nums[i-1]
			})

			nums[i-1], nums[j] = nums[j], nums[i-1]
			return
		}
	}
	sort.Ints(nums)
}

func main() {
	a := []int{3, 2, 1}
	nextPermutation(a)
}
