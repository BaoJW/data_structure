package queue

import "testing"

func TestList_QueueDemo(t *testing.T) {
	q := NewList()
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
		if element == i {
			t.Logf("queue dequeue succeed, element: %d", element)
		}
	}
}
