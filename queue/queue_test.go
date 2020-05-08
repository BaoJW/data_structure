package queue

import (
	"fmt"
	"testing"
)

func TestNewMyQueue(t *testing.T) {
	q := NewMyQueue()

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
