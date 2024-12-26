package hashMapsAndSets

func PairSumUnsorted(target int, nums []int) (output []int) {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, has := hashMap[target-nums[i]]
		if has {
			return append(output, hashMap[target-nums[i]], i)
		}
		hashMap[nums[i]] = i
	}
	return output
}
