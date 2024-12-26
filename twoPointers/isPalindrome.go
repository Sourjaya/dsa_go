package twoPointers

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
func toLowerCase(s string) string {
	result := []rune(s)
	for i := 0; i < len(result); i++ {
		if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] += 32
		}
	}
	return string(result)
}
func IsPalindrome(input string) bool {
	input = toLowerCase(input)
	left, right := 0, len(input)-1
	for left < right {
		for left < right && !isAlphaNumeric(input[left]) {
			left += 1
		}
		for left < right && !isAlphaNumeric(input[right]) {
			right -= 1
		}
		if input[left] != input[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}
