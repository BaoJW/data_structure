package queue

/*
	队列 先入先出的数据结构
    在 FIFO 数据结构中，将首先处理添加到队列中的第一个元素。
	队列是典型的 FIFO 数据结构。
	插入（insert）操作也称作入队（enqueue），新元素始终被添加在队列的末尾。
	删除（delete）操作也被称为出队（dequeue)。 你只能移除第一个元素。
*/

type Element interface{}

type Queue interface {
	InitQueue()                    // 初始化操作，创建一个空队列
	DestroyQueue()                 // 若队列存在，则销毁他
	ClearQueue()                   // 将队列清空
	QueueEmpty() bool              // 若队列为空，返回true，否则返回false
	GetHead() Element              // 若队列存在且非空，用e返回队列Q的队头元素
	EnQueue(element Element) error // 若队列Q存在，插入新元素e到队列Q中并成为队尾元素
	DeQueue() (Element, error)     // 删除队列中Q中队头元素，并用e返回其值
	QueueLength() int              // 返回队列的元素个数
}

type MyQueue struct {
	data         []uint64 // store elements
	startPointer int      // a pointer to indicate the start position
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		data:         make([]uint64, 0, 512), //这里分配的长度以实际需求为准，多次的内存重分配会影响性能，故能一次性分配最好
		startPointer: 0,
	}
}

// Insert an element into the queue. Return true if the operation is successful.
func (q *MyQueue) EnQueue(value uint64) bool {
	q.data = append(q.data, value)
	return true
}

// Delete an element from the queue. Return true if the operation is successful.
func (q *MyQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.startPointer++
	return true
}

// Checks whether the queue is empty or not.
func (q *MyQueue) IsEmpty() bool {
	return q.startPointer >= len(q.data)
}

// Get the front item from the queue.
func (q *MyQueue) Front() uint64 {
	return q.data[q.startPointer]
}

/*
	注意：
	上面的实现很简单，但在某些情况下效率很低。随着起始指针的移动，浪费了越来越多的空间。当我们有空间限制时，这将是难以接受的。
	让我们考虑一种情况，即我们只能分配一个最大长度为5的数组。当我们只添加少于5个元素时，我们的解决方案很有效。例如，如果我们只调用入队函数四次后还想将元素10入队，那么我们可以成功。
	但是我们不能接受更多的入队请求，这是合理的，因为现在队列已经满了。但是如果将一个元素出队呢？实际上，在这种情况下，我们应该还能接受一个元素，但是这块内存空间浪费了，因为我们不会去循环的使用内存空间，该方法只是向下增加空间。
*/
