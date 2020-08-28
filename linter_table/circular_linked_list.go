package linter_table

import "fmt"

// 循环链表 将单链表中终端结点的指针端由空指针改为指向头结点，就使这个单链表行成一个环，这种头尾相接的单链表称为单循环链表
type CircularLinkedList struct {
	Head   *Node
	Length int64
}

func NewCircularLinkedList() LinterTable {
	return &CircularLinkedList{}
}

func (c *CircularLinkedList) Init(length int64) {
	c.Head = new(Node)
}

func (c *CircularLinkedList) IsEmpty() bool {
	if c.Head != nil && c.Head.Next != nil {
		return false
	}

	return true
}

func (c *CircularLinkedList) Clear() {
	c.Head.Next = nil
	c.Length = 0
}

func (c *CircularLinkedList) GetElem(index int64) Element {
	if c.IsEmpty() || index > c.Length || index < 0 {
		fmt.Println("no data")
	}

	p := c.Head
	for i := 0; i < int(index-1); i++ {
		p = p.Next
	}

	return p.Data
}

func (c *CircularLinkedList) Locatelem(e Element) int64 {
	if c.IsEmpty() {
		return 0
	}

	p := c.Head
	var index int64
	for i := 0; i < int(c.Length-1); i++ {
		if p.Data == e {
			index = int64(i)
		}
	}

	return index
}

// 在链表后追加元素
func (c *CircularLinkedList) Append(e Element) {
	node := &Node{Data: e, Next: c.Head}
	p := c.Head

	// 如果链表不为空则遍历
	if !c.IsEmpty() {
		for p.Next != c.Head {
			p = p.Next
		}
	}
	p.Next = node
	c.Length++
	return
}

// 在链表前追加元素
func (c *CircularLinkedList) PreAppend(e Element) {
	node := &Node{Data: e, Next: c.Head}
	c.Head.Next = node
	c.Length++

	return
}

func (c *CircularLinkedList) Insert(index int64, e Element) bool {
	if index < 0 || c.IsEmpty() {
		node := &Node{Data: e, Next: c.Head}
		c.Head.Next = node
		c.Length++
	} else if index > c.Length {
		c.Append(e)
	} else {
		p := c.Head

		for i := 0; i < int(index-1); i++ {
			p = p.Next
		}

		node := &Node{Data: e, Next: nil}

		node.Next = p.Next
		p.Next = node

		c.Length++
	}

	return true
}

func (c *CircularLinkedList) Delete(index int64) bool {
	p := c.Head

	for i := 0; i < int(index-1); i++ {
		p = p.Next
	}

	p.Next = p.Next.Next
	c.Length--
	return true
}

func (c *CircularLinkedList) Len() int64 {
	return c.Length
}

// 打印链表信息
func (c *CircularLinkedList) Print() {
	if c.IsEmpty() {
		fmt.Println(" CircularLinkedList is empty")
		return
	}

	p := c.Head.Next
	i := 1
	for p != c.Head {

		fmt.Printf("iNode：%d, data: %#v\n", i, p.Data)
		i++
		p = p.Next

	}
	fmt.Printf("\n")

}
