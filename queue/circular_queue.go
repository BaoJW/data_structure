package queue

/*
	此前，我们提供了一种简单但低效的队列实现。
	更有效的方法是使用循环队列。具体来说，我们可以使用固定大小的数组和两个指针来指示起始位置和结束位置。目的是重用我们之前提到的被浪费的存储。
*/

/*
	设计循环队列
	循环队列是一种线性数据结构，其操作表现基于FIFO（先进先出）原则并且队尾被连接在队首之后形成一个循环。它也被称为"环形缓冲器"。
	循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通的队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值、
*/

type MyCircularQueue struct {
	data        []int64
	headPointer int
	tailPointer int
	size        int
}

// Initialize your data structure here. Set the size of the queue to be k.
func NewMyCircularQueue(k int) *MyCircularQueue {
	return &MyCircularQueue{
		data:        make([]int64, k, k),
		headPointer: -1,
		tailPointer: -1,
		size:        k,
	}
}

// Insert an element into the circular queue. Return true if the operation is successful.
func (q *MyCircularQueue) EnQueue(value int64) bool {
	if q.IsFull() {
		return false
	}

	if q.IsEmpty() {
		q.headPointer = 0
	}
	q.tailPointer = (q.tailPointer + 1) % q.size
	q.data[q.tailPointer] = value
	return true
}

// Delete an element from the circular queue. Return true if the operation is successful.
func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	if q.headPointer == q.tailPointer {
		q.headPointer = -1
		q.tailPointer = -1
		return true
	}

	q.headPointer = (q.headPointer + 1) % q.size
	return true
}

// Get the front item from the queue.
func (q *MyCircularQueue) Front() int64 {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.headPointer]
}

// Get the last item from the queue.
func (q *MyCircularQueue) Rear() int64 {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.tailPointer]
}

// Checks whether the circular queue is empty or not.
func (q *MyCircularQueue) IsEmpty() bool {
	return q.headPointer == -1
}

// Checks whether the circular queue is full or not.
func (q *MyCircularQueue) IsFull() bool {
	return ((q.tailPointer + 1) % q.size) == q.headPointer
}
