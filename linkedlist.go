package glinq

type LinkedList[T any] struct {
	head, tail *node[T]
	count      int
}

type node[T any] struct {
	val  T
	prev *node[T]
	next *node[T]
}

func NewLinkedListOf[T any](array []T) *LinkedList[T] {
	return nil
}

func (lst *LinkedList[T]) IsEmpty() bool {
	return lst.head == nil && lst.tail == nil
}

func (lst *LinkedList[T]) Push(v T) {
	lst.count++
	n := &node[T]{
		val:  v,
		prev: lst.tail,
		next: nil,
	}
	if lst.IsEmpty() {
		lst.head = n
		lst.tail = n
	}
	lst.tail.next = n
	lst.tail = n
}

func (lst *LinkedList[T]) RemoveAt(index int) {
	panic("not supported")
}

func (lst *LinkedList[T]) Remove(data T) {
	lst.count--
	for p := lst.head; p != nil; p = p.next {
		if Equal(data, p.val) {
			if p == lst.head {
				lst.head = p.next
			}
			if p == lst.tail {
				lst.tail = p.prev
			}
			if p.prev != nil {
				p.prev.next = p.next
			}
			if p.next != nil {
				p.next.prev = p.prev
			}
			return
		}
	}
}

func (lst *LinkedList[T]) Count() int {
	return lst.count
}

func (lst *LinkedList[T]) Contains(elem T) bool {
	containsElem := false
	DoWhile(lst.GetEnumerator(), func(item T) bool {
		if Equal(item, elem) {
			containsElem = true
			return false
		}
		return true
	})
	return containsElem
}

func (lst *LinkedList[T]) GetEnumerator() IEnumerator[T] {
	return &LinkedListEnumerable[T]{first: &lst.head, next: &lst.head}
}

// ToSlice from ArrayList[T] to []T
func (lst *LinkedList[T]) ToSlice() []T {
	slice := make([]T, lst.count)
	index := 0
	DoWhile(lst.GetEnumerator(), func(item T) bool {
		index++
		slice[index] = item
		return false
	})
	return slice
}
