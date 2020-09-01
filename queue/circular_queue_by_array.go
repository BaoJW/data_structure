package queue

import "errors"

type CircularArray struct {
	E     []Element
	front int // 头指针
	rear  int // 尾指针，若队列不空，指向队列尾元素的下一个位置
}

// 循环队列的数组实现
func NewCircularArray() Queue {
	return &CircularArray{}
}

func (c *CircularArray) InitQueue() {
	c.E = make([]Element, 0, MaxSize)
}

func (c *CircularArray) DestroyQueue() {
	c.E = nil
	c = nil
}

func (c *CircularArray) ClearQueue() {
	c.E = c.E[:0]
	c.front = 0
	c.rear = 0
}

func (c *CircularArray) QueueEmpty() bool {
	if c.front == c.rear {
		return true
	}

	return false
}

func (c *CircularArray) GetHead() Element {
	if c.QueueEmpty() {
		return nil
	}

	return c.E[c.front]
}

func (c *CircularArray) EnQueue(element Element) error {
	// 队列满的判断
	if (c.rear+1)%MaxSize == c.front {
		return errors.New("current queue is full")
	}

	c.E = append(c.E, element)
	c.rear = (c.rear + 1) % MaxSize
	return nil
}

func (c *CircularArray) DeQueue() (Element, error) {
	// 队列为空的判断
	if c.front == c.rear {
		return nil, errors.New("current queue is empty")
	}

	tempElement := c.E[c.front]
	c.front = (c.front + 1) % MaxSize
	return tempElement, nil

}

func (c *CircularArray) QueueLength() int {
	return (c.rear - c.front + MaxSize) % MaxSize
}
