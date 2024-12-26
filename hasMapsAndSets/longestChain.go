package hashMapsAndSets

import (
	"github.com/Sourjaya/dsa/dStructures"
)

func LongestChain(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numSet := dStructures.NewHashSet()
	for _, num := range nums {
		numSet[num] = struct{}{}
	}
	longestChain := 0

	for num := range numSet {
		if _, exists := numSet[num-1]; !exists {
			currentNum := num
			currentChain := 1
			for {
				if _, exists := numSet[currentNum+1]; exists {
					currentNum++
					currentChain++
				} else {
					break
				}
			}

			// Update the longest chain found so far
			if currentChain > longestChain {
				longestChain = currentChain
			}
		}
	}
	return longestChain
}
