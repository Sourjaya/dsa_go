package twoPointers

func LargestContainer(heights []int) int {
	left, right := 0, len(heights)-1
	maxWater := 0
	for left < right {
		water := min(heights[left], heights[right]) * (right - left)
		maxWater = max(maxWater, water)
		if heights[left] < heights[right] {
			left += 1
		} else if heights[left] > heights[right] {
			right -= 1
		} else {
			left += 1
			right -= 1
		}
	}
	return maxWater
}
