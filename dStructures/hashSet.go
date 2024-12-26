package dStructures

type Void struct{}

type HashSet map[int]Void

func NewHashSet() HashSet {
	return make(HashSet)
}
