## 队列的定义
```
队列是只允许在一端进行插入操作，而在另一端进行删除操作的线性表
```
队列是一种先进先出(First In First Out)的线性表，简称FIFO。允许插入的一端称为队尾，允许删除的一端称为队头。


## 总结回顾
栈(Stack)是限定仅在表尾进行插入和删除操作的线性表
队列(Queue)是只允许在一端进行插入操作，而在另一端进行删除操作的线性表

它们均可以用线性表的顺序存储结构来实现，但都存在着顺序存储的一些弊端。因此它们各自有各自的技巧来解决这个问题。

对于栈来说，如果是两个相同数据类型的栈，则可以用数组的两端作栈底的方法来让两个栈共享数据，这就可以最大化的利用数组的空间。

对于队列来说，为了避免数组插入和删除时需要移动数据，于是就引入了循环队列，使得队头和队尾可以在数组中循环变化。解决了移动数据的时间损耗，使得未来插入和删除是O(N)的时间复杂度变成了O(1)。