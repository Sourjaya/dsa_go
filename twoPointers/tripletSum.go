package twoPointers

import "sort"

func TripletSum(nums []int) (result [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		pairs := PairSumSortedAllPairs(-nums[i], i+1, nums)
		for j := 0; j < len(pairs); j++ {
			result = append(result, append([]int{nums[i]}, pairs[j]...))
		}
	}
	return result
}
