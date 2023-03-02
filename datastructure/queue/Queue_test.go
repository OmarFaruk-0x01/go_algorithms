package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	q := NewSliceQueue[int](3)
	err := q.Enqueue(20)
	assert.False(t, q.IsEmpty())
	assert.Nil(t, err)

	q.Enqueue(30)
	q.Enqueue(10)
	assert.True(t, q.IsFull())

	err = q.Enqueue(40)
	assert.NotNil(t, err)
}

func TestDequeue(t *testing.T) {
	q := NewSliceQueue[int](3)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(10)

	val, err := q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 20, val)
	q.Dequeue()

	val, err = q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 10, val)
	assert.True(t, q.IsEmpty())

	_, err = q.Dequeue()
	assert.NotNil(t, err)
}
