package linter_table

import "fmt"

type DoubleLinkedNode struct {
	Data Element
	Next *DoubleLinkedNode
	Pre  *DoubleLinkedNode
}

type DoubleLinkedList struct {
	Head   *DoubleLinkedNode
	Length int64
}

// 双向链表
// 在单链表中，有了next指针，就使得我们要查找下一个节点为O(1)，可是如果我们要查找的是上一结点的话，那最坏的时间复杂度就是O(N)
// 双向链表是在单链表的每个结点中，再设置一个指向其前驱结点的指针域
func NewDoubleLinkedList() LinterTable {
	return &DoubleLinkedList{}
}

func (d *DoubleLinkedList) Init(length int64) {
	d.Head = new(DoubleLinkedNode)
	d.Length = 0
}

func (d *DoubleLinkedList) IsEmpty() bool {
	if d.Head != nil && d.Head.Next != nil {
		return false
	}

	return true
}

func (d *DoubleLinkedList) Clear() {
	d.Head.Next = nil
	d.Length = 0
}

func (d *DoubleLinkedList) GetElem(index int64) Element {
	if d.IsEmpty() || index > d.Length || index < 0 {
		fmt.Println("no data")
	}

	p := d.Head
	for i := 0; i < int(index-1); i++ {
		p = p.Next
	}

	return p.Data
}

func (d *DoubleLinkedList) Locatelem(e Element) int64 {
	if d.IsEmpty() {
		return 0
	}

	p := d.Head
	var index int64
	for i := 0; i < int(d.Length-1); i++ {
		if p.Data == e {
			index = int64(i)
		}
	}

	return index
}

// 在链表后追加元素
func (d *DoubleLinkedList) Append(e Element) {
	node := &DoubleLinkedNode{Data: e, Next: nil, Pre: nil}
	p := d.Head

	// 如果链表不为空则遍历
	if !d.IsEmpty() {
		for p.Next != nil {
			p = p.Next
		}
	}

	node.Pre = p
	p.Next = node

	d.Length++
	return
}

// 在链表前追加元素
func (d *DoubleLinkedList) PreAppend(e Element) {
	node := &DoubleLinkedNode{Data: e, Next: nil, Pre: nil}

	node.Next = d.Head.Next
	node.Pre = d.Head

	d.Head.Next = node

	d.Length++

	return
}

func (d *DoubleLinkedList) Insert(index int64, e Element) bool {
	if index < 0 || d.IsEmpty() {
		d.PreAppend(e)
	} else if index > d.Length {
		d.Append(e)
	} else {
		p := d.Head

		for i := 0; i < int(index-1); i++ {
			p = p.Next
		}

		node := &DoubleLinkedNode{Data: e, Next: nil, Pre: nil}

		node.Next = p.Next
		p.Next.Pre = node

		p.Next = node
		node.Pre = p

		d.Length++
	}

	return true
}

func (d *DoubleLinkedList) Delete(index int64) bool {
	p := d.Head

	for i := 0; i < int(index-1); i++ {
		p = p.Next
	}

	p.Pre.Next = p.Next
	p.Next.Pre = p.Pre

	d.Length--
	return true
}

func (d *DoubleLinkedList) Len() int64 {
	return d.Length
}

// 打印链表信息
func (d *DoubleLinkedList) Print() {
	if d.IsEmpty() {
		fmt.Println(" list is empty")
		return
	}

	p := d.Head.Next
	i := 1
	for p != nil {

		fmt.Printf("iNode：%d, data: %#v\n", i, p.Data)
		i++
		p = p.Next

	}
	fmt.Printf("\n")
}
