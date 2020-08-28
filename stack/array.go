package stack

const MaxSize = 10

type Array struct {
	E []Element
}

// 栈的顺序结构（数组）的表现方式
func NewArray() Stack {
	return &Array{}
}

func (a *Array) InitStack() {
	a.E = make([]Element, 0, MaxSize)
}

func (a *Array) DestoryStack() bool {
	a.E = nil
	return true
}

func (a *Array) ClearStack() bool {
	a.E = a.E[:0]
	return true
}

func (a *Array) StackEmpty() bool {
	if len(a.E) <= 0 {
		return true
	}

	return false
}

func (a *Array) GetTop() Element {
	if a.E == nil || len(a.E) <= 0 {
		return nil
	}

	return a.E[len(a.E)-1]
}

func (a *Array) Push(e Element) bool {
	if a.E == nil {
		return false
	}

	// 栈满的情况不在允许插入元素
	if a.Length() >= MaxSize {
		return false
	}

	a.E = append(a.E, e)
	return true
}

func (a *Array) Pop() (Element, bool) {
	if a.E == nil || len(a.E) <= 0 {
		return nil, false
	}

	tempElement := a.E[len(a.E)-1]
	a.E = a.E[0 : len(a.E)-1]
	return tempElement, true
}

func (a *Array) Length() int64 {
	if a.E == nil {
		return 0
	}

	return int64(len(a.E))
}
