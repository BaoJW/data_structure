package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularArray_CircularQueueDemo(t *testing.T) {
	q := NewCircularArray()

	q.InitQueue()

	if q.QueueEmpty() {
		for i := 1; i <= 5; i++ {
			_ = q.EnQueue(i)
		}
	}

	for i := 1; i <= 5; i++ {
		element, _ := q.DeQueue()
		if i == element {
			t.Logf("queue dequeue succeed, element: %d", element)
		}
	}
	_, err := q.DeQueue()
	assert.EqualError(t, err, "current queue is empty", "")

	for i := 1; i <= 10; i++ {
		_ = q.EnQueue(i)
	}

	err = q.EnQueue(6)
	assert.EqualError(t, err, "current queue is full", "")

}
