package linkedlist

import (
	"errors"
	"fmt"
)

const (
	index_out_of_bound = "index out of bound"
	empty_list         = "list is empty"
	remove_item        = "unable to remove item the list is empty"
	unsupported_index  = "index must be a positive number"
)

// Node is a cell of a LinkedList. It stores two data side by side. One is the actual element and Second is the pointer of the next node.
//
// Parameters:
//
//	T: The type of elements stored in the linked list. The type must be comparable.
//
// Fields:
//
//	value: Actual element for the node.
//	next: A pointer to the next node of this node.
type node[T comparable] struct {
	value T
	next  *node[T]
}

// LinkedList is a data structure that stores a sequence of elements. Each element is linked to the next element in the sequence through a "next" pointer. The first element in the sequence is called the "head" of the list, and the last element is called the "tail" of the list.
//
// Parameters:
//
//	T: The type of elements stored in the linked list. The type must be comparable.
//
// Fields:
//
//	first: A pointer to the first element (head) of the linked list.
//	last: A pointer to the last element (tail) of the linked list.
//	size: The number of elements stored in the linked list.
//
// Description:
//
//	The LinkedList type is implemented as a generic struct in Go. The type parameter T specifies the type of elements stored in the linked list. The type parameter must be a comparable type, which means it supports the == and != operators.
//
// Example:
//
//	// Create a new linked list of integers
//	myList := linkedlist.New[int]()
//	// Add elements to the list
//	myList.AddFirst(5)
//	myList.AddLast(10)
//	// Get the first element in the list
//	firstElement := myList.First()
//	// Get the last element in the list
//	lastElement := myList.Last()
type linkedList[T comparable] struct {
	first *node[T]
	last  *node[T]
	size  int
	_nil  T
}

// Create a new Instance of LinkedList with the Given Type
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.addLast(1)
//	myList.addLast(2)
//	myList.addFirst(3)
//	myList.ToSlice() // [3 1 2]
//
//	myList := linkedlist.New[string]()
//	myList.addLast("is")
//	myList.addLast("a")
//	myList.addFirst("This")
//	myList.addLast("LinkedList")
//	myList.ToSlice() // [This is a LinkedList]
func New[T comparable]() linkedList[T] {
	return linkedList[T]{size: 0}
}

// AddFirst adds an element to the beginning of the linked list.
//
// Parameters:
//
//	item: The value of the element to add to the list.
//
// Complexity:
//
//	Time - O(1)
//	Space - O(1)
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.AddFirst(100) // myList: 100
//	myList.AddFirst(102) // myList: 102 -> 100
func (l *linkedList[T]) AddFirst(item T) {
	newNode := &node[T]{value: item}
	if l.isEmpty() {
		l.first = newNode
		l.last = newNode
	} else {
		newNode.next = l.first
		l.first = newNode
	}
	l.size++
}

// AddLast adds an element to the end of the linked list.
//
// Parameters:
//
//	item: The value of the element to add to the list.
//
// Complexity:
//
//	Time - O(1)
//	Space - O(1)
//
// Example:
//
//	myList := linkedlist.New[string]()
//	myList.AddLast("one") // myList: one
//	myList.AddFirst("two") // myList: one -> two
func (l *linkedList[T]) AddLast(item T) {
	newNode := &node[T]{value: item}
	if l.isEmpty() {
		l.first = newNode
		l.last = newNode
	} else {
		oldLastNode := l.last
		l.last = newNode
		oldLastNode.next = l.last
	}
	l.size++
}

// InsertAt adds an element to the given index of the linked list. If the index is bigger then the size of the list It will add the new Item into the end of the list
//
// Parameters:
//
//	item: The value of the element to add to the list.
//	index: The targeted index of the element to add to the list.
//
// Complexity:
//
//	Time - O(1) (Best Case)
//	Time - O(n) (Worst Case)
//	Space - O(1)
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.InsertAt(100, 0) // myList: 100
//	myList.InsertAt(102, 1) // myList: 102 -> 100
func (l *linkedList[T]) InsertAt(item T, index int) {
	switch {
	case index == 0:
		l.AddFirst(item)
		return
	case index >= l.size:
		l.AddLast(item)
		return
	case index > 0 && index < l.size:
		var previousNode *node[T]
		var nextNode *node[T]
		newNode := node[T]{value: item}
		count := 0
		for node := l.first; node != nil; node = node.next {
			if count == index-1 {
				nextNode = node.next
				previousNode = node
				break
			}
			count++
		}
		previousNode.next = &newNode
		newNode.next = nextNode
		l.size++
	}
}

// RemoveFirst remove an element from the beginning of the linked list. If the list is empty it will return a error otherwise nil.
//
// Complexity:
//
//	Time - O(1)
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.InsertAt(100, 0) // myList: 100
//	myList.InsertAt(102, 1) // myList: 102 -> 100
//	_ := myList.RemoveFirst() // myList:  100
func (l *linkedList[T]) RemoveFirst() error {
	if l.isEmpty() {
		return errors.New(remove_item)
	}
	newFirst := l.first.next
	l.first = newFirst
	l.size--
	return nil
}

// RemoveLast remove an element from the end of the linked list. If the list is empty it will return a error otherwise nil.
//
// Complexity:
//
//	Time - O(n)
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.InsertAt(100, 0) // myList: 100
//	myList.InsertAt(102, 1) // myList: 102 -> 100
//	_ := myList.RemoveLast() // myList:  102
func (l *linkedList[T]) RemoveLast() error {
	if l.isEmpty() {
		return errors.New(remove_item)
	}
	count := 0
	for node := l.first; node != nil; node = node.next {
		if count == l.size-2 {
			l.last = node
			l.last.next = nil
			l.size--
			break
		}
		count++
	}
	return nil
}

// RemoveAt remove an element from the given index of the linked list. If the list is empty it will return a error otherwise nil.
//
// Parameters:
//
//	index: The targeted index of the element to remove from the list.
//
// Complexity:
//
//	Time - O(1) (Base Case)
//	Time - O(n) (Worst Case)
//	Space - O(1)
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.InsertFirst(100) // myList: 100
//	myList.InsertFirst(102) // myList: 102 -> 100
//	myList.InsertFirst(103) // myList: 103 -> 102 -> 100
//	myList.RemoveAt(1) // myList:  103 -> 100
func (l *linkedList[T]) RemoveAt(index int) error {
	if l.isEmpty() {
		return errors.New(remove_item)
	}
	if index < 0 {
		return errors.New(unsupported_index)
	}
	if index >= l.size {
		return errors.New(index_out_of_bound)
	}

	switch {
	case index == 0:
		l.RemoveFirst()
		return nil
	case index == l.size-1:
		l.RemoveLast()
		return nil
	case index > 0 && index < l.size:
		var previousNode *node[T]
		var targetedNode *node[T]
		count := 0
		for node := l.first; node != nil; node = node.next {
			if count == index-1 {
				targetedNode = node.next
				previousNode = node
				break
			}
			count++
		}
		previousNode.next = targetedNode.next
		targetedNode.next = nil
		l.size--
	}
	return nil
}

// Find the index an element from the linked list.
//
// Parameters:
//
//	item: A targeted item to find the index from the list
//
// Returns:
//
//	index int: Item Index If the targeted item exist in the list. Otherwise -1
//
// Complexity:
//
//	Time - O(1) (Best Case)
//	Time - O(n) (Worst Case)
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.InsertAt(100, 0) // myList: 100
//	myList.InsertAt(102, 1) // myList: 102 -> 100
//	myList.FindIndex(102) // 1
func (l *linkedList[T]) FindIndex(item T) int {
	index := 0
	for node := l.first; node != nil; node = node.next {
		if node.value == item {
			return index
		}
		index++
	}
	return -1
}

// Traversal all elements of the linked list from the beginning to end.
//
// Parameters:
//
//	function(item T, index int): A anonymous function that call for each item.
//
// Complexity:
//
//	Time - O(n)
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.InsertAt(100, 0) // myList: 100
//	myList.InsertAt(102, 1) // myList: 102 -> 100
//	myList.Traversal(func (item int, index int) {
//		fmt.Println(item) // 102, 100
//	})
func (l linkedList[T]) Traversal(traversal_func func(item T, index int)) {
	index := 0
	for node := l.first; node != nil; node = node.next {
		traversal_func(node.value, index)
		index++
	}
}

// Create a copy of a LinkedList to an slice.
//
// Returns:
//
//	slice []T: A slice contains all node values
//
// Complexity:
//
//	Time - O(n)
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.InsertAt(100, 0) // myList: 100
//	myList.InsertAt(102, 1) // myList: 102 -> 100
//	myList.ToSlice() // myList:  [102 100]
func (l linkedList[T]) ToSlice() []T {
	values := []T{}
	l.Traversal(func(item T, index int) {
		values = append(values, item)
	})
	return values
}

// Get ta total size of the linkedlist.
//
// Returns:
//
//	size int: Total length of the linkedlink.
//
// Complexity:
//
//	Time - O(1)
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.InsertAt(100, 0) // myList: 100
//	myList.InsertAt(102, 1) // myList: 102 -> 100
//	myList.Size() // myList:  2
func (l linkedList[T]) Size() int {
	return l.size
}

// Check the given item exist or not in the linkedlist.
//
// Parameters:
//
//	item: A targeted item to check its existence
//
// Returns:
//
//	bool: If targeted item exist return true otherwise false.
//
// Complexity:
//
//	Time - O(n)
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.InsertAt(100, 0) // myList: 100
//	myList.InsertAt(102, 1) // myList: 102 -> 100
//	myList.Contains(102) // true
//	myList.Contains(103) // false
func (l linkedList[T]) Contains(item T) bool {
	for node := l.first; node != nil; node = node.next {
		if node.value == item {
			return true
		}
	}
	return false
}

// Get the first item of the linkedlist.
//
// Returns:
//
//	item: The first item of the linkedlist.
//
// Complexity:
//
//	Time - O(1)
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.InsertAt(100, 0) // myList: 100
//	myList.InsertAt(102, 1) // myList: 102 -> 100
//	myList.First() // 102
func (l linkedList[T]) First() (T, error) {
	if l.last == nil {
		return l._nil, errors.New(empty_list)
	}

	return l.first.value, nil
}

// Get the last item of the linkedlist.
//
// Returns:
//
//	item: The last item of the linkedlist.
//
// Complexity:
//
//	Time - O(1)
//
// Example:
//
//	myList := linkedlist.New[int]()
//	myList.InsertAt(100, 0) // myList: 100
//	myList.InsertAt(102, 1) // myList: 102 -> 100
//	myList.Last() // 100
func (l linkedList[T]) Last() (T, error) {
	if l.last == nil {
		return l._nil, errors.New(empty_list)
	}

	return l.last.value, nil
}

func (l linkedList[T]) isEmpty() bool {
	return l.size == 0
}

// Overwrite the Stringer Interface Function For this Struct
func (l linkedList[T]) String() string {
	return fmt.Sprintf("%v", l.ToSlice())
}
