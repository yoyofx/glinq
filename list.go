package glinq

import (
	"errors"
	"reflect"
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

// Add adds an element to the list
func (list *ArrayList[T]) Add(x T) { list.data = append([]T(list.data), x) }

// RemoveAt removes an element from the list.
func (list *ArrayList[T]) RemoveAt(i int) {
	if i < len([]T(list.data)) {
		list.data = append([]T(list.data[:i]), list.data[i+1:]...)
	} else {
		panic(errors.New("array index out of bounds"))
	}
}

// Contains return bool,that element is in the list.
func (list *ArrayList[T]) Contains(elem T) bool {
	for _, item := range list.data {
		if reflect.ValueOf(item).Interface() == reflect.ValueOf(elem).Interface() {
			return true
		}
	}
	return false
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
