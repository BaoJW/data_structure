package queue

import (
	"fmt"
	"testing"
)

func TestNewMyCircularQueue(t *testing.T) {
	q := NewMyCircularQueue(5)

	q.EnQueue(5)
	q.EnQueue(3)
	if !q.IsEmpty() {
		fmt.Println(q.Front())
	}

	q.DeQueue()
	if !q.IsEmpty() {
		fmt.Println(q.Front())
	}

	q.DeQueue()
	if !q.IsEmpty() {
		fmt.Println(q.Front())
	}
}
