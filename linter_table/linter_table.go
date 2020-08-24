package linter_table

// 线性表的定义： 零个或多个数据元素的有限序列

type Element interface{}

// 线性表的抽象数据类型 这里展示了几个比较基本的方法，提供了一个demo,具体的要根据实际的业务需要灵活变化使用
type LinterTable interface {
	Init(length int64)                  // 初始化操作，建立一个空的线性表, 入参是线性表的容量
	IsEmpty() bool                      // 若线性表为空，返回true，否则返回false
	Clear()                             // 将线性表清空
	GetElem(index int64) Element        // 返回线性表中第i个元素的的值
	Locatelem(e Element) int64          // 在线性表中查找给定的元素，如果查找成功，返回该元素在表中的序号表示成功；否则，返回0表示失败
	Insert(index int64, e Element) bool // 在线性表中的第i个位置插入新元素e
	Delete(index int64) bool            // 删除线性表中第i个位置的元素
	Len() int64                         // 返回线性表L的元素个数
	Print()                             // 打印信息
}
