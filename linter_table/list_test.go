package linter_table

import (
	"fmt"
	"testing"
)

func TestList_ListModel(t *testing.T) {
	l := NewList()
	l.Init(10)
	l.Insert(2, 222)
	l.Insert(4, 444)
	l.Insert(3, 333)
	l.Insert(3, 555)

	l.Delete(2)

	fmt.Println(l.GetElem(2))

	l.Print()
	fmt.Println(l.Len())

	l.IsEmpty()

	l.Print()
	fmt.Println(l.Len())
}
