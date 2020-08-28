package stack

type Element interface{}

type Stack interface {
	InitStack()           // 初始化操作，建立一个空栈
	DestoryStack() bool   // 若栈存在，则销毁他，销毁成功返回true
	ClearStack() bool     // 将栈清空，清空成功返回true
	StackEmpty() bool     // 如栈为空，返回true
	GetTop() Element      // 获取栈顶的元素
	Push(e Element) bool  // 压栈，成功返回true
	Pop() (Element, bool) // 弹栈，成功返回true，并返回弹出的元素
	Length() int64        // 获取栈空间元素个数
}
