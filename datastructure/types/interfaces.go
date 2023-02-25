package types

type LinkedList[T comparable] interface {
	AddFirst(item T)
	AddLast(item T)
	InsertAt(item T, index int)
	RemoveAt(index int) error
	RemoveFirst() error
	RemoveLast() error
	Traversal(func(item T, index int))
	FindIndex(item T) int
	ToSlice() []T
	Contains(item T) bool
	Size() int
	First() (T, error)
	Last() (T, error)
}

type StackQueue[T comparable] interface {
	Push(item T)
	Pop() (T, error)
	Peek() (T, error)
	Size() int
	IsEmpty() bool
}
