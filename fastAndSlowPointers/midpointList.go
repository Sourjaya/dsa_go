package fastAndSlowPointers

import "github.com/Sourjaya/dsa/dStructures"

func MidpointList(list dStructures.LinkedList[int]) *dStructures.Node[int] {
	slow, fast := list.Head, list.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
