package dStructures

import "fmt"

// Node represents a single node in the linked list.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// LinkedList represents the entire linked list structure.
type LinkedList[T comparable] struct {
	Head *Node[T]
}

func NewNode[T comparable](value T) *Node[T] {
	return &Node[T]{Value: value}
}

// Add adds a new node with the given value to the end of the list.
func (l *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

// Print prints the linked list values.
func (l *LinkedList[T]) Print() {
	current := l.Head
	for current != nil && current.Next != nil {
		fmt.Print(current.Value, "-->")
		current = current.Next
	}
	fmt.Print(current.Value)
	fmt.Println()
}

// Size returns the number of elements in the linked list.
func (l *LinkedList[T]) Size() int {
	count := 0
	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

// Remove removes the first occurrence of the value from the linked list.
func (l *LinkedList[T]) Remove(value T) {
	if l.Head == nil {
		return
	}

	// If the head node contains the value
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}

	// Traverse the list and remove the node
	current := l.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}
}
