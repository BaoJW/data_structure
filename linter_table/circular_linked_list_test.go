package linter_table

import "testing"

func TestCircularLinkedList_CircularLinkedListDemo(t *testing.T) {
	circularLinkedList := NewCircularLinkedList()

	circularLinkedList.Init(0)

	circularLinkedList.Insert(1, 10)
	circularLinkedList.Print()

	circularLinkedList.Insert(2, 20)
	circularLinkedList.Insert(3, 30)
	circularLinkedList.Insert(5, 50)
	circularLinkedList.Print()

	circularLinkedList.Insert(4, 40)
	circularLinkedList.Print()
}
