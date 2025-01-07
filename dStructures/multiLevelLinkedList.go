package dStructures

import (
	"fmt"
)

type MultiLevelListNode[T comparable] struct {
	Val   T
	Next  *MultiLevelListNode[T]
	Child *MultiLevelListNode[T]
}

// NewMultiLevelListNode creates and returns a new MultiLevelListNode.
func NewMultiLevelListNode[T comparable](val T) *MultiLevelListNode[T] {
	return &MultiLevelListNode[T]{Val: val}
}

// Add appends a new node with a value to the current level of the multilevel linked list.
func Add[T comparable](head *MultiLevelListNode[T], val T) *MultiLevelListNode[T] {
	newNode := NewMultiLevelListNode(val)
	if head == nil {
		return newNode
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return head
}

// Print recursively prints the multilevel linked list, showing multiple levels clearly.
func Print[T comparable](head *MultiLevelListNode[T], level int) {
	if head == nil {
		return
	}

	current := head
	for current != nil {
		fmt.Printf("Level %d: %v\n", level, current.Val)
		if current.Child != nil {
			Print(current.Child, level+1)
		}
		current = current.Next
	}
}
