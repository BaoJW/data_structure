package stack

import (
	"fmt"
	"testing"
)

func TestArray_ArrayDemo(t *testing.T) {
	a := NewArray()
	a.InitStack()

	a.Push(1)
	a.Push(2)
	a.Push(3)
	a.Push(4)
	a.Push(5)
	a.Push(6)
	a.Push(7)

	if a.GetTop() == 7 {
		fmt.Println("stack push success")
	}

	popElement, _ := a.Pop()
	if popElement == 7 && a.GetTop() == 6 {
		fmt.Println("stack pop success")
	}

}
