package hashMapsAndSets

func GeometricTriplets(r int, nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	leftMap := make(map[int]int)
	rightMap := make(map[int]int)

	for _, num := range nums {
		rightMap[num]++
	}

	count := 0

	for _, num := range nums {
		rightMap[num]--

		if num%r == 0 {
			left := num / r
			right := num * r

			if leftFreq, leftExists := leftMap[left]; leftExists {
				if rightFreq, rightExists := rightMap[right]; rightExists {
					count += leftFreq * rightFreq
				}
			}
		}

		leftMap[num]++
	}

	return count
}
