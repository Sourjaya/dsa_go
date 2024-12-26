// Problem Statement
/*
 Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number.

 Examples:

Input: arr[] = {2, 7, 11, 15}, target = 9
Output: 1 2
Explanation: Since the array is 1-indexed, arr[1] + arr[2] = 2 + 7 =  9

Input: {1, 3, 4, 6, 8, 11} target = 10
Output: 3 4
Explanation: Since the array is 1-indexed, arr[3] + arr[5] = 4 + 6 = 10
*/

package twoPointers

func PairSumSortedAllPairs(target int, index int, nums []int) (result [][]int) {
	var pair []int
	left, right := index, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			left += 1
		} else if sum > target {
			right -= 1
		} else {
			result = append(result, append(pair, nums[left], nums[right]))
			left += 1
			for left < right && nums[left] == nums[left-1] {
				left += 1
			}
		}
	}
	return result
}

func PairSumSorted(target int, nums []int) (result []int) {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			left += 1
		} else if sum > target {
			right -= 1
		} else {
			result = append(result, left, right)
			break
		}
	}
	return result
}
