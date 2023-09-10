package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCacheScenarios(t *testing.T) {

	checkList := func(cache *lruCache, expectedForwardList []KeyPair) {
		list := cache.list.ToList()
		reverseList := cache.list.GetReverseList()

		var j int
		var expectedReverseList = make([]KeyPair, len(expectedForwardList))

		for i := len(expectedForwardList) - 1; i >= 0; i-- {
			expectedReverseList[j] = expectedForwardList[i]
			j++
		}

		if !reflect.DeepEqual(list, expectedForwardList) {
			t.Errorf("unexpected list after Put(1, 1), got: %v, want: %v", list, expectedForwardList)
		}

		if !reflect.DeepEqual(reverseList, expectedReverseList) {
			t.Errorf("unexpected reverse list after Put(1, 1). got: %v, want: %v", reverseList, expectedReverseList)
		}

	}

	cache, _ := NewLRUCache(3)

	cache.Put(1, 1)
	checkList(cache, []KeyPair{{1, 1}})

	cache.Put(2, 2)
	checkList(cache, []KeyPair{{2, 2}, {1, 1}})

	cache.Put(3, 3)
	checkList(cache, []KeyPair{{3, 3}, {2, 2}, {1, 1}})

	got := cache.Get(0)
	if got != -1 {
		t.Errorf("got a value in the cache that doesn't exist")
	}

	// check cache eviction of value 4
	cache.Put(4, 4)
	checkList(cache, []KeyPair{{4, 4}, {3, 3}, {2, 2}})

	// check update of cache
	cache.Put(2, 5)
	checkList(cache, []KeyPair{{2, 5}, {4, 4}, {3, 3}})

	node := cache.Get(3)

	if node != 3 {
		t.Error("failed to get correct value")
	}

	checkList(cache, []KeyPair{{3, 3}, {2, 5}, {4, 4}})

	cache.Put(1, 1)
	checkList(cache, []KeyPair{{1, 1}, {3, 3}, {2, 5}})

}

func TestNewLRUCache(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name   string
		args   args
		want   *lruCache
		expErr error
	}{
		{
			name: "invalid capacity",
			args: args{
				capacity: -2,
			},
			want:   nil,
			expErr: errInvalidCapacity,
		},
		{
			name: "makes cache with 4 capacity",
			args: args{
				capacity: 4,
			},
			want: &lruCache{
				capacity: 4,
			},
			expErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewLRUCache(tt.args.capacity)
			if err != tt.expErr {
				t.Errorf("NewLRUCache() error = %v, wantErr %v", err, tt.expErr)
				return
			}
		})
	}
}

func Test_lruCache_Get(t *testing.T) {

	type args struct {
		key int
	}

	tests := []struct {
		name      string
		initCache func() *lruCache
		args      args
		want      int
	}{
		{
			name: "key doesnt exist in cache",
			initCache: func() *lruCache {
				cache, err := NewLRUCache(5)
				if err != nil {
					t.Error("failed to create cache ", err)
				}
				return cache
			},
			args: args{
				key: 1,
			},
			want: -1,
		},
		{
			name: "gets key-value from the cache with given key",
			initCache: func() *lruCache {
				cache, err := NewLRUCache(5)
				if err != nil {
					t.Error("failed to create cache ", err)
				}

				cache.Put(0, 5)
				cache.Put(1, 10)
				cache.Put(2, 15)

				return cache
			},
			args: args{
				key: 1,
			},
			want: 10,
		},
		{
			name: "updates the key-value to most recently used",
			initCache: func() *lruCache {
				cache, err := NewLRUCache(5)
				if err != nil {
					t.Error("failed to create cache ", err)
				}

				cache.Put(1, 1)
				cache.Put(2, 0)
				cache.Put(5, 5)

				return cache
			},
			args: args{
				key: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initCache()
			got := c.Get(tt.args.key)

			if got != tt.want {
				t.Errorf("lruCache.Get() = %v, want %v", got, tt.want)
			}

			head := c.list.Head()

			if got != -1 && head != nil && head.Val.Key != tt.args.key {
				t.Error("cache didnt update key to most recently used")
			}

		})
	}
}

func Test_lruCache_Put(t *testing.T) {

	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name              string
		initCache         func() *lruCache
		args              args
		expectedKeyPairs  []KeyPair
		expectedCacheSize int
		expectedCacheKeys []int
	}{
		{
			name: "Puts a new value into empty cache",
			initCache: func() *lruCache {
				cache, err := NewLRUCache(1)
				if err != nil {
					t.Error("failed to init cache, ", err)
				}
				return cache
			},
			args: args{
				key:   1,
				value: 2,
			},
			expectedKeyPairs:  []KeyPair{{Key: 1, Val: 2}},
			expectedCacheSize: 1,
			expectedCacheKeys: []int{1},
		},
		{
			name: "Updates existing key with value in the cache",
			initCache: func() *lruCache {

				cache, _ := NewLRUCache(5)

				cache.Put(1, 0)
				cache.Put(10, 4)
				return cache
			},
			args: args{
				key:   1,
				value: 2,
			},
			expectedKeyPairs:  []KeyPair{{Key: 1, Val: 2}, {Key: 10, Val: 4}},
			expectedCacheSize: 2,
			expectedCacheKeys: []int{1, 10},
		},
		{
			name: "Adds a new key-val pair to a cache with max capacity, since cache is at max capacity it will remove the least recently used",
			initCache: func() *lruCache {
				cache, err := NewLRUCache(1)

				if err != nil {
					t.Error("failed to init cache, ", err)
				}

				// fills up the cache
				cache.list.Insert(KeyPair{Key: 1, Val: 1})
				return cache
			},
			args: args{
				key:   3,
				value: 6,
			},
			expectedKeyPairs:  []KeyPair{{Key: 3, Val: 6}},
			expectedCacheSize: 1,
			expectedCacheKeys: []int{3},
		},
		{
			name: "updates an existing node in a max capacity cache",
			initCache: func() *lruCache {
				cache, err := NewLRUCache(3)

				node1 := DoubleLinkedListNode{
					Val: KeyPair{Key: 1, Val: 1},
				}

				node2 := DoubleLinkedListNode{
					Val: KeyPair{Key: 2, Val: 5},
				}

				node3 := DoubleLinkedListNode{
					Val: KeyPair{Key: 6, Val: 9},
				}

				cache.list.InsertAtHead(&node3)
				cache.list.InsertAtHead(&node2)
				cache.list.InsertAtHead(&node1)

				// 1 -> 2 -> 6

				cache.values = map[int]*DoubleLinkedListNode{
					1: &node1,
					2: &node2,
					3: &node3,
				}

				if err != nil {
					t.Error("failed to init cache, ", err)
				}

				return cache
			},
			args: args{
				key:   2,
				value: 3,
			},
			expectedKeyPairs:  []KeyPair{{Key: 2, Val: 3}, {Key: 1, Val: 1}, {Key: 6, Val: 9}},
			expectedCacheSize: 3,
			expectedCacheKeys: []int{2, 1, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initCache()
			c.Put(tt.args.key, tt.args.value)

			if c.list.length != tt.expectedCacheSize {
				t.Errorf("cache size is unexpected. got: %v, want: %v", c.list.length, tt.expectedCacheSize)
			}

			// TODO: check keys are correct (ideally we would like to check nodes as well but I'm not sure how to do considering pointers)
			keys := make([]int, len(c.values))
			i := 0
			for k := range c.values {
				keys[i] = k
				i++
			}
			fmt.Println(keys)

			keyPairs := c.list.ToList()

			if !reflect.DeepEqual(keyPairs, tt.expectedKeyPairs) {
				t.Errorf("expected values: %v, but got: %v", tt.expectedKeyPairs, keyPairs)
			}
		})
	}
}

func TestDoubleLinkedList_Insert(t *testing.T) {
	type fields struct {
		head   func() DoubleLinkedListNode
		length int
	}
	type args struct {
		value KeyPair
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedLength   int
		expectedElements []KeyPair
	}{
		{
			name: "inserts a new keyvalue pair into an empty list",
			fields: fields{
				head: func() DoubleLinkedListNode {
					return NewList().head
				},
				length: NewList().length,
			},
			args: args{
				value: KeyPair{Key: 1, Val: 2},
			},
			expectedLength: 1,
			expectedElements: []KeyPair{
				{Key: 1, Val: 2},
			},
		},
		{
			name: "add to existing list",
			fields: fields{
				length: 1,
				head: func() DoubleLinkedListNode {
					l := NewList()

					l.Insert(KeyPair{1, 1})

					return l.head
				},
			},
			args: args{
				value: KeyPair{2, 2},
			},
			expectedLength:   2,
			expectedElements: []KeyPair{{1, 1}, {2, 2}},
		},
		{
			name: "daddy spank me",
			fields: fields{
				length: 2,
				head: func() DoubleLinkedListNode {
					l := NewList()

					l.Insert(KeyPair{1, 1})
					l.Insert(KeyPair{2, 2})

					// 1 -> 2
					return l.head
				},
			},
			args: args{
				value: KeyPair{Key: 3, Val: 3},
			},
			expectedLength: 3,
			expectedElements: []KeyPair{
				{Key: 1, Val: 1},
				{Key: 2, Val: 2},
				{Key: 3, Val: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DoubleLinkedList{
				head:   tt.fields.head(),
				length: tt.fields.length,
			}
			l.Insert(tt.args.value)

			// expect that new value is at end of list

			keyPairs := l.ToList()

			reverseList := l.GetReverseList()

			var forwardsList = make([]KeyPair, len(reverseList))

			var j int
			for i := len(reverseList) - 1; i >= 0; i-- {
				forwardsList[j] = reverseList[i]
				j++
			}

			lastKeyPair := keyPairs[len(keyPairs)-1]

			if lastKeyPair != tt.args.value {
				t.Errorf("last key pair are not equal! got: %v, want: %v", lastKeyPair, tt.args.value)
			}

			if !reflect.DeepEqual(keyPairs, tt.expectedElements) {
				t.Errorf("key pairs are not equal! got: %v, want: %v", keyPairs, tt.expectedElements)
			}

			if !reflect.DeepEqual(keyPairs, forwardsList) {
				// this error suggests that the tail pointers aren't being setup correctly
				t.Errorf("forward list does not equal the same list in reverse. keyPairs: %v, reverseList: %v, forwardsList: %v", keyPairs, reverseList, forwardsList)
			}

			// expect head.Prev to point to new value

			tail := l.head.Prev

			if tail.Val != tt.args.value {
				t.Errorf("tail value is incorrect! got: %v, want: %v", tail.Val, tt.args.value)
			}

			// expect length

			if tt.expectedLength != l.length {
				t.Errorf("length is not equal got: %v, want: %v", l.length, tt.expectedLength)
			}

		})
	}
}

func TestDoubleLinkedList_Remove(t *testing.T) {

	node := DoubleLinkedListNode{
		Val: KeyPair{Key: 1, Val: 2},
	}

	node7 := DoubleLinkedListNode{
		Val: KeyPair{Key: 7, Val: 7},
	}

	setup := func() DoubleLinkedListNode {
		var node1, node2 *DoubleLinkedListNode

		head := NewList().head

		node2 = &DoubleLinkedListNode{
			Next: &head,
			Val:  KeyPair{Key: 2, Val: 6},
		}

		node1 = &DoubleLinkedListNode{
			Prev: &head,
			Next: node2,
			Val:  KeyPair{Key: 1, Val: 7},
		}

		node2.Prev = node1

		head.Next = node1
		head.Prev = node2

		return head
	}

	type fields struct {
		head   func() DoubleLinkedListNode
		length int
	}
	type args struct {
		node *DoubleLinkedListNode
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedLength   int
		expectedKeyPairs []KeyPair
	}{
		{
			name: "removes from an empty list",
			fields: fields{
				head: func() DoubleLinkedListNode {
					return NewList().head
				},
				length: NewList().length,
			},
			args: args{
				node: &DoubleLinkedListNode{
					Val: KeyPair{Key: 1, Val: 1},
				},
			},
			expectedLength:   0,
			expectedKeyPairs: []KeyPair{},
		},
		{
			name: "removes 1st and only node from the list",
			fields: fields{
				head: func() DoubleLinkedListNode {
					head := NewList().head
					head.Next = &node
					head.Prev = &node
					return head
				},
				length: 1,
			},
			expectedLength:   0,
			expectedKeyPairs: []KeyPair{},
			args: args{
				node: &node,
			},
		},
		{
			name: "removes node from list",
			fields: fields{
				head:   setup,
				length: 2,
			},
			args: args{
				node: setup().Next,
			},
			expectedLength:   1,
			expectedKeyPairs: []KeyPair{{Key: 2, Val: 6}},
		},
		{
			name: "removes middle node from list",
			fields: fields{
				head: func() DoubleLinkedListNode {
					list := NewList()

					list.Insert(KeyPair{Key: 5, Val: 5})

					// 5 -> 2
					list.Insert(KeyPair{Key: 2, Val: 2})

					// 1 -> 5 -> 2
					list.InsertAtHead(&node)

					// 7 -> 1 -> 5 -> 2
					list.InsertAtHead(&node7)

					return list.head
				},
				length: 4,
			},
			args: args{
				// remove 1 so result should be: 7 -> 5 -> 2
				node: &node,
			},
			expectedLength:   3,
			expectedKeyPairs: []KeyPair{{Key: 7, Val: 7}, {Key: 5, Val: 5}, {Key: 2, Val: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DoubleLinkedList{
				head:   tt.fields.head(),
				length: tt.fields.length,
			}
			l.Remove(tt.args.node)

			if tt.expectedLength != l.length {
				t.Errorf("length after remove is not equal! got: %v, want: %v", l.length, tt.expectedLength)
			}

			keyPairs := l.ToList()

			if !reflect.DeepEqual(keyPairs, tt.expectedKeyPairs) {
				t.Errorf("key pairs not equal! got: %v, want: %v", keyPairs, tt.expectedKeyPairs)
			}

		})
	}
}

func TestDoubleLinkedList_InsertAtHead(t *testing.T) {
	type args struct {
		node *DoubleLinkedListNode
	}
	tests := []struct {
		name             string
		setup            func() *DoubleLinkedList
		args             args
		expectedLength   int
		expectedKeyPairs []KeyPair
	}{
		{
			name:  "inserts at head of empty list",
			setup: NewList,
			args: args{
				node: &DoubleLinkedListNode{
					Val: KeyPair{Key: 1, Val: 1},
				},
			},
			expectedLength:   1,
			expectedKeyPairs: []KeyPair{{Key: 1, Val: 1}},
		},
		{
			name: "inserts into list with some values",
			setup: func() *DoubleLinkedList {
				l := NewList()
				l.Insert(KeyPair{Key: 2, Val: 2})
				return l
			},
			args: args{
				node: &DoubleLinkedListNode{
					Val: KeyPair{Key: 1, Val: 1},
				},
			},
			expectedLength:   2,
			expectedKeyPairs: []KeyPair{{Key: 1, Val: 1}, {Key: 2, Val: 2}},
		},
		{
			name: "the ultimate daddy test",
			setup: func() *DoubleLinkedList {
				l := NewList()
				l.Insert(KeyPair{Key: 4, Val: 4})
				l.Insert(KeyPair{Key: 3, Val: 3})
				l.Insert(KeyPair{Key: 2, Val: 2})
				l.Insert(KeyPair{Key: 1, Val: 1})

				// 4 -> 3 -> 2 -> 1

				return l
			},
			args: args{
				node: &DoubleLinkedListNode{
					Val: KeyPair{Key: 5, Val: 5},
				},
			},
			expectedLength:   5,
			expectedKeyPairs: []KeyPair{{Key: 5, Val: 5}, {Key: 4, Val: 4}, {Key: 3, Val: 3}, {Key: 2, Val: 2}, {Key: 1, Val: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.setup()
			l.InsertAtHead(tt.args.node)

			keyPairs := l.ToList()

			reverseList := l.GetReverseList()

			var forwardsList = make([]KeyPair, len(reverseList))

			var j int
			for i := len(reverseList) - 1; i >= 0; i-- {
				forwardsList[j] = reverseList[i]
				j++
			}

			if !reflect.DeepEqual(keyPairs, forwardsList) {
				// this error suggests that the tail pointers aren't being setup correctly
				t.Errorf("forward list does not equal the same list in reverse. keyPairs: %v, reverseList: %v, forwardsList: %v", keyPairs, reverseList, forwardsList)
			}

			if !reflect.DeepEqual(keyPairs, tt.expectedKeyPairs) {
				t.Errorf("key pairs not the same! got: %v, want: %v", keyPairs, tt.expectedKeyPairs)
			}

			if l.length != tt.expectedLength {
				t.Errorf("length not the same! got: %v, want: %v", l.length, tt.expectedLength)
			}

		})
	}
}
