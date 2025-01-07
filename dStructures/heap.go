package dStructures

import "fmt"

/* Heap type with generics */
type Heap[T any] struct {
	Array      []T
	Comparator func(parent, child T) bool
}

/* Constructor function */
func NewHeap[T any](comparator func(parent, child T) bool) *Heap[T] {
	return &Heap[T]{
		Array:      []T{},
		Comparator: comparator,
	}
}

/* find parent index */
func parent(index int) int {
	return (index - 1) / 2
}

/* find left Child index */
func leftChild(index int) int {
	return (2 * index) + 1
}

/* find right Child index */
func rightChild(index int) int {
	return (2 * index) + 2
}

func (h *Heap[T]) swap(i, j int) {
	h.Array[i], h.Array[j] = h.Array[j], h.Array[i]
}

func (h *Heap[T]) Insert(key T) {
	h.Array = append(h.Array, key)
	/* Heapify up to the root */
	h.heapifyUp(len(h.Array) - 1)
}

func (h *Heap[T]) heapifyUp(index int) {

	/* until root is reached or the heap property is not satisfied continue to loop */
	for index > 0 && !h.Comparator(h.Array[parent(index)], h.Array[index]) {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *Heap[T]) heapifyDown(index int) {
	lastIndex := len(h.Array) - 1
	left, right := leftChild(index), rightChild(index)
	childToCompare := 0

	/* Continue while index has at least one child */
	for left <= lastIndex {
		if right <= lastIndex && !h.Comparator(h.Array[left], h.Array[right]) {
			childToCompare = right
		} else {
			childToCompare = left
		}

		/* If the current node satisfies the heap property, stop */
		if h.Comparator(h.Array[index], h.Array[childToCompare]) {
			break
		}

		/* Swap the current node with the selected child */
		h.swap(index, childToCompare)
		index = childToCompare
		left, right = leftChild(index), rightChild(index)
	}
}

func (h *Heap[T]) Delete() (T, error) {
	if len(h.Array) == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("heap is empty")
	}

	/* Get the root element */
	root := h.Array[0]

	/* Move the last element to the root and remove the last element */
	h.Array[0] = h.Array[len(h.Array)-1]
	h.Array = h.Array[:len(h.Array)-1]

	/* Heapify down from the root */
	h.heapifyDown(0)

	return root, nil
}
