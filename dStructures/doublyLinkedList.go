package dStructures

import "fmt"

// DoublyNode represents a single DoublyNode in the doubly linked list.
// DoublyNode represents a single node in the doubly linked list.
type DoublyNode[T comparable] struct {
	Key   T
	Value T
	Next  *DoublyNode[T]
	Prev  *DoublyNode[T]
}

// DoublyLinkedList represents the entire doubly linked list structure.
type DoublyLinkedList[T comparable] struct {
	Head *DoublyNode[T]
	Tail *DoublyNode[T]
}

// NewNode creates and returns a new doubly linked list node.
func NewDoublyNode[T comparable](key T, value T) *DoublyNode[T] {
	return &DoublyNode[T]{Key: key, Value: value}
}

// Add adds a new node with the given key and value to the end of the list.
func (l *DoublyLinkedList[T]) Add(key T, value T) {
	newNode := &DoublyNode[T]{Key: key, Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Prev = l.Tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

// Print prints the doubly linked list values in forward order.
func (l *DoublyLinkedList[T]) Print() {
	current := l.Head
	for current != nil {
		if current.Next != nil {
			fmt.Printf("(%v, %v)<-->", current.Key, current.Value)
		} else {
			fmt.Printf("(%v, %v)", current.Key, current.Value)
		}
		current = current.Next
	}
	fmt.Println()
}

// PrintReverse prints the doubly linked list values in reverse order.
func (l *DoublyLinkedList[T]) PrintReverse() {
	current := l.Tail
	for current != nil {
		if current.Prev != nil {
			fmt.Printf("(%v, %v)<-->", current.Key, current.Value)
		} else {
			fmt.Printf("(%v, %v)", current.Key, current.Value)
		}
		current = current.Prev
	}
	fmt.Println()
}

// Size returns the number of elements in the doubly linked list.
func (l *DoublyLinkedList[T]) Size() int {
	count := 0
	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

// Remove removes the first occurrence of the key from the doubly linked list.
func (l *DoublyLinkedList[T]) Remove(key T) {
	if l.Head == nil {
		return
	}

	// If the head node contains the key
	if l.Head.Key == key {
		if l.Head == l.Tail {
			l.Head = nil
			l.Tail = nil
		} else {
			l.Head = l.Head.Next
			l.Head.Prev = nil
		}
		return
	}

	current := l.Head
	for current != nil && current.Key != key {
		current = current.Next
	}

	if current != nil {
		if current == l.Tail {
			l.Tail = current.Prev
			l.Tail.Next = nil
		} else {
			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev
		}
	}
}
