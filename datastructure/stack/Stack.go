package stack

import (
	"errors"
	"fmt"
)

const (
	stack_empty_error = "stack is empty"
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
		return s._nil, errors.New(stack_empty_error)
	}
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	s.size--
	return top, nil
}

func (s *stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return s._nil, errors.New(stack_empty_error)
	}
	return s.stack[len(s.stack)-1], nil
}

func (s stack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s stack[T]) Size() int {
	return s.size
}

func (s stack[T]) String() string {
	return fmt.Sprintf("%v", s.stack)
}
