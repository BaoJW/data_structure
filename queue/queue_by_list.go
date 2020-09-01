package queue

type Nodes struct {
	data Element
	next *Nodes
}

type List struct {
	front *Nodes // 头指针
	rear  *Nodes // 尾指针
	Len   int
}

func NewList() Queue {
	return &List{}
}

func (l *List) InitQueue() {
	l.front = &Nodes{data: nil, next: nil}
	l.rear = &Nodes{data: nil, next: nil}
}

func (l *List) DestroyQueue() {
	l.front = nil
	l.rear = nil
	l = nil
}

func (l *List) ClearQueue() {
	l.front = nil
	l.rear = nil
}

func (l *List) QueueEmpty() bool {
	if l.front.next == nil {
		return true
	}

	return false
}

func (l *List) GetHead() Element {
	return l.front.data
}

func (l *List) EnQueue(element Element) error {
	node := &Nodes{data: element, next: nil}

	if l.QueueEmpty() {
		l.front = node
		l.rear = l.front
	}

	l.rear.next = node
	l.rear = node
	l.Len++
	return nil

}

func (l *List) DeQueue() (Element, error) {
	tempNode := l.front
	l.front = l.front.next
	l.Len--
	return tempNode.data, nil

}

func (l *List) QueueLength() int {
	return l.Len
}
