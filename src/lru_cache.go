package main

import "errors"

/**


LRU = Least Recently Used

We are making a LRU cache which stores a list of items up to the given cache capacity.

The cache has two main operations:

Get(key int) -> Gets a value in the cache by it's key. If the value doesn't exist we return -1 (for the sake of this shitty leetcode test). When we get the given value we have to update it so it's at the front of the cache (as it's now the most recently used)
Put(key int, value int) -> If the item doesn't already exist in the cache then we add the key-value pair. If the number of keys exceeds the capacity then we have to evict the least recently used key. If the item already exists then update it.



*/

type lruCache struct {
	capacity int
}

var errInvalidCapacity = errors.New("capacity has to be a positive number")

func NewLRUCache(capacity int) (*lruCache, error) {
	if capacity <= 0 {
		return nil, errInvalidCapacity
	}

	return &lruCache{
		capacity: capacity,
	}, nil
}

// Get - gets the value of a cache by it's key. If the key doesn't exist we return -1
func (c *lruCache) Get(key int) int {
	return -1
}

// If the item doesn't already exist in the cache then we add the key-value pair. If the number of keys exceeds the capacity then we have to evict the least recently used key. If the item already exists then update it.
func (c *lruCache) Put(key int, value int) {

}
