package linkedlist

import (
	"github.com/Sourjaya/dsa/dStructures"
)

func RemoveKthLastNode(list *dStructures.LinkedList[int], k int) *dStructures.LinkedList[int] {
	dummy := dStructures.NewNode(0)
	dummy.Next = list.Head
	head := list.Head
	trailer, leader := dummy, dummy
	for i := 0; i < k; i++ {

		leader = leader.Next
		if leader == nil {
			return &dStructures.LinkedList[int]{Head: head}
		}
	}

	for leader.Next != nil {
		leader = leader.Next
		trailer = trailer.Next
	}

	trailer.Next = trailer.Next.Next
	return &dStructures.LinkedList[int]{Head: dummy.Next}
}
