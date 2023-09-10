package main

import (
	"errors"
)

/**


LRU = Least Recently Used

We are making a LRU cache which stores a list of items up to the given cache capacity.

The cache has two main operations:

Get(key int) -> Gets a value in the cache by it's key. If the value doesn't exist we return -1 (for the sake of this shitty leetcode test). When we get the given value we have to update it so it's at the front of the cache (as it's now the most recently used)
Put(key int, value int) -> If the item doesn't already exist in the cache then we add the key-value pair. If the number of keys exceeds the capacity then we have to evict the least recently used key. If the item already exists then update it.

Get and Put must be O(1) time complexity

*/

type KeyPair struct {
	Key int
	Val int
}

type DoubleLinkedListNode struct {
	Val  KeyPair
	Next *DoubleLinkedListNode
	Prev *DoubleLinkedListNode
}

type DoubleLinkedList struct {
	head   DoubleLinkedListNode
	length int
}

// Inits / clears list
func (l *DoubleLinkedList) Init() *DoubleLinkedList {
	l.head.Next = &l.head
	l.head.Prev = &l.head
	l.length = 0
	return l
}

// Makes a new list ready for use
func NewList() *DoubleLinkedList {
	return new(DoubleLinkedList).Init()
}

// Inserts element at the end of the list
func (l *DoubleLinkedList) Insert(value KeyPair) {

	tail := l.Tail()

	node := DoubleLinkedListNode{
		Val:  value,
		Next: &l.head,
		Prev: tail,
	}

	// head points to new tail
	l.head.Prev = &node

	// previous tail now points to node
	tail.Next = &node

	// we only have one value so it points to itself
	if l.head.Val.Key == l.head.Next.Val.Key {
		l.head.Next = &node
	}

	l.length = l.length + 1
}

// returns the head node
func (l *DoubleLinkedList) Head() *DoubleLinkedListNode {
	if l == nil || l.head.Next == nil {
		return nil
	}
	return l.head.Next
}

// returns the tail node
func (l *DoubleLinkedList) Tail() *DoubleLinkedListNode {
	if l == nil || l.head.Prev == nil {
		return nil
	}
	return l.head.Prev
}

func (l *DoubleLinkedList) InsertAtHead(node *DoubleLinkedListNode) {
	firstNode := l.Head()

	// update the existing node
	if firstNode.Val.Key == node.Val.Key {
		firstNode.Next = node.Next
		firstNode.Prev = node.Prev
		firstNode.Val = node.Val
		return
	}

	node.Prev = &l.head

	// update the oldHead to point to our new node instead of the dummy head node
	firstNode.Prev = node
	node.Next = firstNode

	l.head.Next = node

	if l.length == 0 {
		// 1st node being inserted so head == tail
		l.head.Prev = node
	}

	l.length++

}

// Removes a node from the linked list
func (l *DoubleLinkedList) Remove(node *DoubleLinkedListNode) {
	if node == nil || l.length == 0 {
		return
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	// head is not a pointer so we have to manually do it here
	if l.head.Next != nil {
		if l.head.Next.Val.Key == node.Val.Key {
			l.head.Next = node.Next
		}
	}

	if l.head.Prev != nil {
		if l.head.Prev.Val.Key == node.Val.Key {
			l.head.Prev = node.Prev
		}
	}

	// *delete it by assigning as nil
	node = nil

	l.length--
}

func (l *DoubleLinkedList) FindNode(key int) *DoubleLinkedListNode {
	return l.head.findNode(key, 0, l.length)

}

// returns a node with matching key or nil
func (l *DoubleLinkedListNode) findNode(key int, counter int, max int) *DoubleLinkedListNode {

	if l.Val.Key == key || l == nil {
		return l
	}

	if counter >= max {
		return nil
	}

	if l.Next != nil {
		return l.Next.findNode(key, counter+1, max)
	}

	return nil

}

const forwards = "forwards"
const backwards = "backwards"

// returns all the KeyPairs in the list
func (l *DoubleLinkedListNode) travese(arr []KeyPair, head DoubleLinkedListNode, counter int, maxSize int, direction string) []KeyPair {

	if counter >= maxSize {
		return arr
	}

	// we've encounted a cycle - assuming here that the keys are unique
	if l.Val.Key == head.Val.Key {
		return arr
	}

	arr = append(arr, l.Val)
	if l.Next != nil && direction == forwards {
		return l.Next.travese(arr, head, counter+1, maxSize, direction)
	}

	if l.Prev != nil && direction == backwards {
		return l.Prev.travese(arr, head, counter+1, maxSize, direction)
	}

	return arr
}

// ToList converts linked list to an array of its values
func (l *DoubleLinkedList) ToList() []KeyPair {
	var values []KeyPair
	if l == nil || l.length == 0 || l.head.Next == nil {
		return []KeyPair{}
	}
	return l.head.Next.travese(values, l.head, 0, l.length, forwards)
}

// GetReverseList converts linked list into array of its values but going backwards from tail -> head
func (l *DoubleLinkedList) GetReverseList() []KeyPair {
	var values []KeyPair
	if l == nil || l.length == 0 || l.head.Prev == nil {
		return []KeyPair{}
	}
	return l.head.Prev.travese(values, l.head, 0, l.length, backwards)
}

type lruCache struct {
	capacity int
	values   map[int]*DoubleLinkedListNode // map of key to pointer of nodes
	list     DoubleLinkedList              // actual list
}

var errInvalidCapacity = errors.New("capacity has to be a positive number")

func NewLRUCache(capacity int) (*lruCache, error) {
	if capacity <= 0 {
		return nil, errInvalidCapacity
	}

	return &lruCache{
		capacity: capacity,
		values:   make(map[int]*DoubleLinkedListNode, capacity),
		list:     *NewList(),
	}, nil
}

// leetcode shit
type LRUCache = lruCache

// this is for leetcode use
func Constructor(capacity int) LRUCache {
	return lruCache{
		capacity: capacity,
		values:   make(map[int]*DoubleLinkedListNode, capacity),
		list:     *NewList(),
	}
}

// Get - gets the value of a cache by it's key. If the key doesn't exist we return -1
// When we get the given value we have to update it so it's at the front of the cache (as it's now the most recently used)
func (c *lruCache) Get(key int) int {
	node, ok := c.values[key]

	if !ok {
		return -1
	}

	// Get can be described as simply Removing the current element then doing an insertion at the head

	c.list.Remove(node)
	c.list.InsertAtHead(node)

	return node.Val.Val
}

// If the item doesn't already exist in the cache then we add the key-value pair.
// internally when we put a value it's inserted at the head of the stack
// If the number of keys exceeds the capacity then we have to evict the least recently used key.
// If the item already exists then update it.
func (c *lruCache) Put(key int, value int) {

	// Put is slightly more difficult theres a few scenarios to consider:

	// 1. cache is not full and data doesn't already exist in cache
	// we create a new node and simply insert at head

	// 2. cache is not full and data already exists in cache
	// we remove the current item then insert it at the head

	// 3. cache is full and data doesn't already exist in cache
	// we remove the tail then we create a new node to insert at head

	// 4. cache is full and data already exists in cache
	// we remove current item and re-insert at the head

	node, exists := c.values[key]

	if c.list.length >= c.capacity && !exists {
		// cache is full
		tail := c.list.Tail()
		c.list.Remove(tail)
		delete(c.values, tail.Val.Key) // remove the reference to the node
	}

	if exists {
		c.list.Remove(node)
		node.Val.Val = value
		c.list.InsertAtHead(node)
	} else {
		// node doesn't exist
		newNode := DoubleLinkedListNode{
			Val: KeyPair{Key: key, Val: value},
		}
		c.list.InsertAtHead(&newNode)
		c.values[key] = &newNode
	}

}
