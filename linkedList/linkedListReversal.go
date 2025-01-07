package linkedlist

import "github.com/Sourjaya/dsa/dStructures"

func LinkedListReversal(list dStructures.LinkedList[int]) *dStructures.LinkedList[int] {
	var prevNode *dStructures.Node[int] = nil
	currNode := list.Head
	for currNode != nil {
		nextNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}
	return &dStructures.LinkedList[int]{Head: prevNode}
}

func LinkedListReversalRecursive(list dStructures.LinkedList[int], head *dStructures.Node[int]) (*dStructures.LinkedList[int], *dStructures.Node[int]) {
	if head == nil || head.Next == nil {
		return &dStructures.LinkedList[int]{Head: head}, head
	}
	_, newHead := LinkedListReversalRecursive(list, head.Next)
	head.Next.Next = head
	head.Next = nil
	return &dStructures.LinkedList[int]{Head: newHead}, newHead
}

func LinkedListReversalUsingHead(head *dStructures.Node[int]) *dStructures.Node[int] {
	var prevNode *dStructures.Node[int] = nil
	currNode := head
	for currNode != nil {
		nextNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}
	return prevNode
}
