package queue

import (
	"testing"
)

func TestArray_QueueDemo(t *testing.T) {
	q := NewArray()
	q.InitQueue()
	if q.QueueEmpty() {
		_ = q.EnQueue(1)
		_ = q.EnQueue(2)
		_ = q.EnQueue(3)
		_ = q.EnQueue(4)
		_ = q.EnQueue(5)
	}

	for i := 1; i <= 5; i++ {
		element, _ := q.DeQueue()
		if i == element {
			t.Log("queue dequeue succeed")
		}
	}

}
