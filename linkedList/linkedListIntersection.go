package linkedlist

import "github.com/Sourjaya/dsa/dStructures"

func LinkedListIntersection(listA, listB dStructures.LinkedList[int]) *dStructures.Node[int] {
	ptrA, ptrB := listA.Head, listB.Head

	for ptrA != ptrB {
		if ptrA != nil {
			ptrA = ptrA.Next
		} else {
			ptrA = listB.Head
		}
		if ptrB != nil {
			ptrB = ptrB.Next
		} else {
			ptrB = listA.Head
		}
	}
	return ptrA
}
