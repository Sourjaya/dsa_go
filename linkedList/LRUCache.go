package linkedlist

import "github.com/Sourjaya/dsa/dStructures"

type LRUCache struct {
	capacity int
	hashMap  map[int]*dStructures.DoublyNode[int]
	head     *dStructures.DoublyNode[int]
	tail     *dStructures.DoublyNode[int]
}

func Constructor(capacity int) *LRUCache {
	head := dStructures.NewDoublyNode(-1, -1)
	tail := dStructures.NewDoublyNode(-1, -1)
	head.Next = tail
	tail.Prev = head
	return &LRUCache{
		capacity: capacity,
		hashMap:  make(map[int]*dStructures.DoublyNode[int]),
		head:     head,
		tail:     tail,
	}
}

func (cache *LRUCache) addToTail(node *dStructures.DoublyNode[int]) {
	prev := cache.tail.Prev
	prev.Next = node
	cache.tail.Prev = node
	node.Prev = prev
	node.Next = cache.tail
}

func (cache *LRUCache) removeNode(node *dStructures.DoublyNode[int]) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (cache *LRUCache) Get(key int) int {
	if value, exists := cache.hashMap[key]; !exists {
		return -1
	} else {
		cache.removeNode(value)
		cache.addToTail(value)
	}
	return cache.hashMap[key].Value
}

func (cache *LRUCache) Put(key int, value int) {
	if value, exists := cache.hashMap[key]; exists {
		cache.removeNode(value)
	}
	node := dStructures.NewDoublyNode(key, value)
	cache.hashMap[key] = node
	if len(cache.hashMap) > cache.capacity {
		delete(cache.hashMap, cache.head.Next.Key)
		cache.removeNode(cache.head.Next)
	}
	cache.addToTail(node)
}
