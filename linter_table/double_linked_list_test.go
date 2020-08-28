package linter_table

import "testing"

func TestDoubleLinkedList_DoubleLinkedListDemo(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()

	doubleLinkedList.Init(0)

	doubleLinkedList.Insert(1, 10)
	doubleLinkedList.Print()

	doubleLinkedList.Insert(2, 20)
	doubleLinkedList.Insert(3, 30)
	doubleLinkedList.Insert(5, 50)
	doubleLinkedList.Print()

	doubleLinkedList.Insert(4, 40)
	doubleLinkedList.Print()
}
