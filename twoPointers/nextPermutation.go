package twoPointers

func reverse(s []rune) {

	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}
func NextPermutation(input []rune) string {
	pivot := len(input) - 2
	for pivot >= 0 && input[pivot] >= input[pivot+1] {
		pivot--
	}
	if pivot == -1 {
		reverse(input)
		return string(input)
	}
	right := len(input) - 1
	for input[right] <= input[pivot] {
		right--
	}
	input[pivot], input[right] = input[right], input[pivot]
	reverse(input[pivot+1:])
	return string(input)
}
