package stack

import "errors"

const (
	StackEmptyError = "stack is empty"
)

type stack[T comparable] struct {
	stack []T
	size  int
	_nil  T
}

func New[T comparable]() stack[T] {
	return stack[T]{stack: []T{}, size: 0}
}

func (s *stack[T]) Push(item T) {
	s.stack = append(s.stack, item)
	s.size++
}

func (s *stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return s._nil, errors.New(StackEmptyError)
	}
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	s.size--
	return top, nil
}

func (s *stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return s._nil, errors.New(StackEmptyError)
	}
	return s.stack[len(s.stack)-1], nil
}

func (s stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s stack[T]) Size() int {
	return s.size
}
