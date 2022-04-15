package glinq

import (
	"errors"
)

// ArrayList Array List for the slice
type ArrayList[T any] struct {
	data []T
}

// NewListOf new list of slice
func NewListOf[T any](array []T) *ArrayList[T] {
	return &ArrayList[T]{data: array}
}

// NewList new list of empty slice
func NewList[T any]() *ArrayList[T] {
	var array []T
	return &ArrayList[T]{data: array}
}

// Push adds an element to the list
func (list *ArrayList[T]) Push(x T) { list.data = append([]T(list.data), x) }

// RemoveAt removes an element from the list by index.
func (list *ArrayList[T]) RemoveAt(i int) {
	if i < len([]T(list.data)) {
		list.data = append([]T(list.data[:i]), list.data[i+1:]...)
	} else {
		panic(errors.New("array index out of bounds"))
	}
}

// Remove removes an element from the list.
func (list *ArrayList[T]) Remove(item T) {
	removeAtIndex := 0
	for index, elem := range list.data {
		if Equal(elem, item) {
			removeAtIndex = index
			break
		}
	}
	list.RemoveAt(removeAtIndex)
}

// Contains return bool,that element is in the list.
func (list *ArrayList[T]) Contains(elem T) bool {
	for _, item := range list.data {
		if Equal(item, elem) {
			return true
		}
	}
	return false
}

func (list *ArrayList[T]) IndexOf(T) int {
	return 0
}

// Count return a number, that's list elements count.
func (list *ArrayList[T]) Count() int {
	return len([]T(list.data))
}

// ToSlice from ArrayList[T] to []T
func (list *ArrayList[T]) ToSlice() []T {
	return list.data
}

// ToQueryable from ArrayList[T] to Queryable[T
func (list *ArrayList[T]) ToQueryable() Queryable[T] {
	return list.data
}

// GetEnumerator get enumerable object
func (list *ArrayList[T]) GetEnumerator() IEnumerator[T] {
	return &ListEnumerable[T]{list: list.data, index: 0}
}
