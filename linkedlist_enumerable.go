package glinq

// LinkedListEnumerable 支持遍历链表元素。
type LinkedListEnumerable[T any] struct {
	first **node[T]
	next  **node[T]
}

// Next 移动迭代器到下一个元互。
// 如果已经到达尾部则返回 false。
func (it *LinkedListEnumerable[T]) Next() bool {
	if *it.next == nil {
		return false
	}
	it.next = &(*it.next).next
	return true
}

// Value 返回当前元素内容。
// 如果元素为空 bool 值为 false。
func (it *LinkedListEnumerable[T]) Value() (T, bool) {
	if *it.next == nil {
		var zero T
		return zero, false
	}
	return (*it.next).val, true
}

func (it *LinkedListEnumerable[T]) Reset() {
	it.next = it.first
}
