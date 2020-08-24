package linter_table

// 线性表的顺序存储结构可以通过一维数组来实现
type Array struct {
	data []Element
}

// 线性表的顺序存储结构的优缺点
// 优点：1. 无须为表中元素之间的逻辑关系而增加额外的存储空间
//		2. 可以快速地存取表中任一位置的元素
// 缺点：1. 插入和删除操作需要移动大量元素（O(N)的时间复杂度）
//		2. 当线性表长度变化较大时，难以确定存储空间的容量
// 		3. 造成存储空间的"碎片"
func NewArray() LinterTable {
	return &Array{}
}

func (a *Array) Init(length int64) {
	a.data = make([]Element, 0, length)
}

func (a *Array) IsEmpty() bool {
	if len(a.data) <= 0 {
		return true
	}

	return false
}

func (a *Array) Clear() {
	a.data = a.data[:0]
}

func (a *Array) GetElem(index int64) Element {
	if index < 0 || int(index) >= len(a.data) {
		return 0
	}
	return a.data[index]
}

func (a *Array) Locatelem(e Element) int64 {
	for index, value := range a.data {
		if value == e {
			return int64(index)
		}
	}

	return 0
}

func (a *Array) Insert(index int64, e Element) bool {
	if index < 0 || int(index) > len(a.data) {
		return false
	}

	if index == 0 {
		a.data = append([]Element{e}, a.data...)
	} else {
		a.data = append(append(a.data[:index], e), a.data[index:]...)
	}

	return true
}

func (a *Array) Delete(index int64) bool {
	if len(a.data) <= 0 || int(index) >= len(a.data) {
		return false
	}
	a.data = append(a.data[:index], a.data[index+1:]...)
	return true
}

func (a *Array) Len() int64 {
	return int64(len(a.data))
}

// 打印数组信息
func (a *Array) Print() {
}
