package linter_table

import "fmt"

// 每一个节点都会有一个数据和指向下一个节点的指针
type Node struct {
	Data Element
	Next *Node
}

// 每条单向链表都有一个头指针和链表长度，当链表中只包含头节点，则链表长度为0，称为空链表。 当插入元素时，也是忽略头节点进行的操作。
type List struct {
	Head   *Node
	Length int64
}

func NewList() LinterTable {
	return &List{}
}

func (l *List) Init(length int64) {
	l.Head = new(Node)
	l.Length = 0
}

func (l *List) IsEmpty() bool {
	if l.Head != nil && l.Head.Next != nil {
		return false
	}

	return true
}

func (l *List) Clear() {
	l.Head.Next = nil
	l.Length = 0
}

func (l *List) GetElem(index int64) Element {
	if l.IsEmpty() || index > l.Length || index < 0 {
		fmt.Println("no data")
	}

	p := l.Head
	for i := 0; i < int(index-1); i++ {
		p = p.Next
	}

	return p.Data
}

func (l *List) Locatelem(e Element) int64 {
	if l.IsEmpty() {
		return 0
	}

	p := l.Head
	var index int64
	for i := 0; i < int(l.Length-1); i++ {
		if p.Data == e {
			index = int64(i)
		}
	}

	return index
}

// 在链表后追加元素
func (l *List) Append(e Element) {
	node := &Node{Data: e, Next: nil}
	p := l.Head

	// 如果链表不为空则遍历
	if !l.IsEmpty() {
		for p.Next != nil {
			p = p.Next
		}
	}
	p.Next = node
	l.Length++
	return
}

// 在链表前追加元素
func (l *List) PreAppend(e Element) {
	node := &Node{Data: e, Next: nil}

	node.Next = l.Head.Next
	l.Head.Next = node

	l.Length++

	return
}

func (l *List) Insert(index int64, e Element) bool {
	if index < 0 || l.IsEmpty() {
		l.PreAppend(e)
	} else if index > l.Length {
		l.Append(e)
	} else {
		p := l.Head

		for i := 0; i < int(index-1); i++ {
			p = p.Next
		}

		node := &Node{Data: e, Next: nil}

		node.Next = p.Next
		p.Next = node

		l.Length++
	}

	return true
}

func (l *List) Delete(index int64) bool {
	p := l.Head

	for i := 0; i < int(index-1); i++ {
		p = p.Next
	}

	p.Next = p.Next.Next
	l.Length--
	return true
}

func (l *List) Len() int64 {
	return l.Length
}

// 打印链表信息
func (l *List) Print() {
	if l.IsEmpty() {
		fmt.Println(" list is empty")
		return
	}

	p := l.Head.Next
	i := 1
	for p != nil {

		fmt.Printf("iNode：%d, data: %#v\n", i, p.Data)
		i++
		p = p.Next

	}
}
