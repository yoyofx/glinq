package glinq

// LinkedList is double linked list data structure
type LinkedList[T any] struct {
	head, tail *node[T]
	count      int
}

// double linked list of node structure
type node[T any] struct {
	val  T
	prev *node[T]
	next *node[T]
}

// NewLinkedListOf creates a new linked list of array.
func NewLinkedListOf[T any](array []T) *LinkedList[T] {
	list := &LinkedList[T]{}
	for _, item := range array {
		list.Push(item)
	}
	return list
}

// NewLinkedList creates a new linked list of empty.
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// IsEmpty list is empty?
func (lst *LinkedList[T]) IsEmpty() bool {
	return lst.head == nil && lst.tail == nil
}

// Push adds an element to the list
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

// RemoveAt removes an element from the list by index.
func (lst *LinkedList[T]) RemoveAt(index int) {
	panic("not supported")
}

// Remove removes an element from the list.
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

// Count return a number, that's list elements count.
func (lst *LinkedList[T]) Count() int {
	return lst.count
}

// Contains return bool,that element is in the list.
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

func (lst *LinkedList[T]) IndexOf(elem T) int {
	index := 0
	DoWhile(lst.GetEnumerator(), func(item T) bool {
		if Equal(elem, item) {
			return false
		}
		index++
		return true
	})
	return index
}

// GetEnumerator get enumerable object
func (lst *LinkedList[T]) GetEnumerator() IEnumerator[T] {
	return &LinkedListEnumerable[T]{first: &lst.head, next: &lst.head}
}

// ToQueryable from LinkedList[T] to Queryable[T
func (lst *LinkedList[T]) ToQueryable() Queryable[T] {
	return lst.ToSlice()
}

// ToSlice from ArrayList[T] to []T
func (lst *LinkedList[T]) ToSlice() []T {
	slice := make([]T, lst.count)
	index := 0
	DoWhile(lst.GetEnumerator(), func(item T) bool {
		slice[index] = item
		index++
		return true
	})
	return slice
}
