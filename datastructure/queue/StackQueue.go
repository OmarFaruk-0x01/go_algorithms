package queue

import (
	"errors"
	"fmt"

	"github.com/OmarFaruk-0x01/go_algorithms/datastructure/stack"
)

type stack_queue[T comparable] struct {
	stack_1 stack.Stack[T]
	stack_2 stack.Stack[T]
	max     int
}

func NewStackQueue[T comparable](maxItem int) *stack_queue[T] {
	stack1 := stack.New[T]()
	stack2 := stack.New[T]()
	return &stack_queue[T]{stack_1: stack1, stack_2: stack2, max: maxItem}
}

func (q *stack_queue[T]) Enqueue(item T) error {
	if q.stack_1.Size() > q.max || q.stack_2.Size() > q.max {
		return errors.New(queue_full_error)
	}
	q.stack_1.Push(item)
	return nil
}

func (q *stack_queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var nil T
		return nil, errors.New(queue_empty_error)
	}
	if q.stack_2.IsEmpty() {
		moveStackData(q.stack_1, q.stack_2)
	}
	item, err := q.stack_2.Pop()
	return item, err
}

func (q *stack_queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var nil T
		return nil, errors.New(queue_empty_error)
	}
	if q.stack_2.IsEmpty() {
		moveStackData(q.stack_1, q.stack_2)
	}
	item, err := q.stack_2.Peek()
	return item, err
}

func (q *stack_queue[T]) IsEmpty() bool {
	return q.stack_1.IsEmpty() && q.stack_2.IsEmpty()
}

func (q *stack_queue[T]) Size() int {
	return q.stack_1.Size() + q.stack_2.Size()
}

func (q *stack_queue[T]) SetMax(max int) {
	q.max = max
}

func (q stack_queue[T]) String() string {
	return fmt.Sprintf("stack1: %v stack2: %v", q.stack_1, q.stack_2)
}

func moveStackData[T comparable](stack1 stack.Stack[T], stack2 stack.Stack[T]) {
	for !stack1.IsEmpty() {
		item, _ := stack1.Pop()
		stack2.Push(item)
	}
}
