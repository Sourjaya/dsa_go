package linkedlist

import "github.com/Sourjaya/dsa/dStructures"

func FlattenList(head *dStructures.MultiLevelListNode[int]) *dStructures.MultiLevelListNode[int] {
	if head == nil {
		return nil
	}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	curr := head
	for curr != nil {
		if curr.Child != nil {
			tail.Next = curr.Child
			curr.Child = nil
			for tail.Next != nil {
				tail = tail.Next
			}
		}
		curr = curr.Next
	}
	return head
}
