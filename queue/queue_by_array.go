package queue

import "errors"

const MaxSize = 10

type Array struct {
	E []Element
}

func NewArray() Queue {
	return &Array{}
}

func (a *Array) InitQueue() {
	a.E = make([]Element, 0, MaxSize)
}

func (a *Array) DestroyQueue() {
	a.E = nil
	a = nil
}

func (a *Array) ClearQueue() {
	a.E = a.E[:0]
}

func (a *Array) QueueEmpty() bool {
	if len(a.E) == 0 {
		return true
	}

	return false
}

func (a *Array) GetHead() Element {
	if len(a.E) <= 0 {
		return nil
	}

	return a.E[0]
}

func (a *Array) EnQueue(element Element) error {
	if len(a.E) >= MaxSize {
		return errors.New("current queue is full")
	}

	a.E = append(a.E, element)
	return nil
}

func (a *Array) DeQueue() (Element, error) {
	if len(a.E) <= 0 {
		return nil, errors.New("current queue is empty")
	}

	tempElement := a.E[0]
	a.E = a.E[1:]
	return tempElement, nil
}

func (a *Array) QueueLength() int {
	return len(a.E)
}
