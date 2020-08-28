package stack

import (
	"fmt"
	"testing"
)

func TestList_ListDemo(t *testing.T) {
	l := NewList()
	l.InitStack()

	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	l.Push(6)
	l.Push(7)

	if l.GetTop() == 7 {
		fmt.Println("stack push success")
	}

	popElement, _ := l.Pop()
	if popElement == 7 && l.GetTop() == 6 {
		fmt.Println("stack pop success")
	}
}
