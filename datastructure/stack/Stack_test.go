package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	stack := New[int]()
	assert.Equal(t, 0, stack.Size())
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
	stack.Push(2)
	assert.Equal(t, 2, stack.Size())
	stack.Push(100)
	assert.Equal(t, 3, stack.Size())
	stack.Push(200)
	assert.Equal(t, 4, stack.Size())
}

func TestPeek(t *testing.T) {
	stack := New[int]()
	val, err := stack.Peek()
	assert.Equal(t, 0, val)
	assert.NotNil(t, err)
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
	stack.Push(2)
	assert.Equal(t, 2, stack.Size())
	stack.Push(100)
	assert.Equal(t, 3, stack.Size())
	stack.Push(200)
	assert.Equal(t, 4, stack.Size())

	val, err = stack.Peek()
	assert.Nil(t, err)
	assert.NotNil(t, val)
	assert.Equal(t, 200, val)

}

func TestPop(t *testing.T) {
	stack := New[int]()
	val, err := stack.Pop()
	assert.Equal(t, 0, val)
	assert.NotNil(t, err)
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
	stack.Push(2)
	assert.Equal(t, 2, stack.Size())

	val, err = stack.Pop()
	assert.Equal(t, 2, val)
	assert.Nil(t, err)

	val, err = stack.Peek()
	assert.Equal(t, 1, val)
	assert.Nil(t, err)

	stack.Push(100)
	assert.Equal(t, 2, stack.Size())
	stack.Push(200)
	assert.Equal(t, 3, stack.Size())

	val, err = stack.Peek()
	assert.Equal(t, 200, val)
	assert.Nil(t, err)

	val, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 200, val)

}
