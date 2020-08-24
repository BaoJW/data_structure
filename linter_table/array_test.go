package linter_table

import "testing"

func TestArray_ArrayModel(t *testing.T) {
	a := NewArray()

	a.Init(10)

	if a.IsEmpty() {
		a.Insert(-1, 10) // 预期插入失败
		a.Insert(0, 10)  // 预期插入成功
		a.Insert(1, 20)  // 预期插入成功
		a.Insert(3, 10)  // 预期插入失败
	}

	if a.Len() == 2 {
		if 10 == a.GetElem(0) && 20 == a.GetElem(1) {
			t.Log("succeed")
		}
	}

	a.Clear()

	if a.IsEmpty() {
		t.Log("array is empty")
	}

}
