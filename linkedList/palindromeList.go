package linkedlist

import (
	"github.com/Sourjaya/dsa/dStructures"
	"github.com/Sourjaya/dsa/fastAndSlowPointers"
)

func IsPalindrome(list dStructures.LinkedList[int]) bool {
	mid := fastAndSlowPointers.MidpointList(list)
	secondHead := LinkedListReversalUsingHead(mid)
	ptr1, ptr2 := list.Head, secondHead
	for ptr2 != nil {
		if ptr1.Value != ptr2.Value {
			return false
		}
		ptr1, ptr2 = ptr1.Next, ptr2.Next
	}
	return true
}
