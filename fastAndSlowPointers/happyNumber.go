package fastAndSlowPointers

func getNextNum(n int) int {
	next := 0
	for n > 0 {
		d := n % 10
		next += d * d
		n = n / 10
	}
	return next
}

func HappyNumber(num int) bool {
	slow, fast := num, num
	for {
		slow = getNextNum(slow)
		fast = getNextNum(getNextNum(fast))
		if fast == 1 {
			return true
		}
		if fast == slow {
			return false
		}
	}
}
