package fastAndSlowPointers

import "github.com/Sourjaya/dsa/dStructures"

func CycleDetection(list dStructures.LinkedList[int]) bool {
	slow, fast := list.Head, list.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
