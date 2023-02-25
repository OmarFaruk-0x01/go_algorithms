package queue

import (
	"errors"
	"fmt"
)

const (
	queue_empty_error = "queue is empty"
	queue_full_error  = "queue is full"
)

type queue[T comparable] struct {
	queue []T
	size  int
	max   int
	_nil  T
}

func New[T comparable](maxItem int) queue[T] {
	return queue[T]{queue: []T{}, size: 0, max: maxItem}
}

func (s *queue[T]) Enqueue(item T) error {
	if s.IsFull() {
		return errors.New(queue_full_error)
	}
	s.queue = append(s.queue, item)
	s.size++
	return nil
}

func (s *queue[T]) Dequeue() (T, error) {
	if s.IsEmpty() {
		return s._nil, errors.New(queue_empty_error)
	}
	top := s.queue[0]
	s.queue = s.queue[1:len(s.queue)]
	s.size--
	return top, nil
}

func (s *queue[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return s._nil, errors.New(queue_empty_error)
	}
	return s.queue[0], nil
}

func (s queue[T]) IsEmpty() bool {
	return s.size == 0
}

func (s queue[T]) Size() int {
	return s.size
}

func (s queue[T]) IsFull() bool {
	return s.max == s.size
}

func (s *queue[T]) SetMax(max int) {
	s.max = max
}

func (s *queue[T]) ToSlice() []T {
	return s.queue
}

func (s queue[T]) String() string {
	return fmt.Sprintf("%v", s.queue)
}
