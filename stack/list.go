package stack

type Node struct {
	Data Element
	Next *Node
}

type List struct {
	Head *Node
	Len  int64
}

// 栈的链式表达方式（链表）
func NewList() Stack {
	return &List{}
}

func (l *List) InitStack() {
	l.Head = new(Node)
}

func (l *List) DestoryStack() bool {
	l.Head = nil // 这里只是将头指针给置为空了，其配置的内存由go的gc去实现
	return true
}

func (l *List) ClearStack() bool {
	l.Head.Next = nil
	return true
}

func (l *List) StackEmpty() bool {
	if l.Head.Next == nil {
		return true
	}

	return false
}

func (l *List) GetTop() Element {
	if l.Head == nil || l.Head.Next == nil {
		return nil
	}

	p := l.Head
	for p.Next != nil {
		p = p.Next
	}

	return p.Data
}

func (l *List) Push(e Element) bool {
	if l.Head == nil {
		return false
	}

	p := l.Head
	for p.Next != nil {
		p = p.Next
	}

	node := &Node{Data: e, Next: nil}
	p.Next = node
	l.Len++

	return true
}

func (l *List) Pop() (Element, bool) {
	if l.Head == nil || l.Head.Next == nil {
		return nil, false
	}

	p := l.Head
	var tempElement Element
	for i := 0; i <= int(l.Len); i++ {
		if i == int(l.Len-1) {
			continue
		}

		if i == int(l.Len) {
			tempElement = p.Next.Data
			break
		}

		p = p.Next
	}

	p.Next = nil
	l.Len--
	return tempElement, true
}

func (l *List) Length() int64 {
	return l.Len
}
