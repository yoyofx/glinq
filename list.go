package glinq

import (
	"errors"
	"reflect"
)

type ArrayList[T any] struct {
	data []T
}

func NewListOf[T any](array []T) *ArrayList[T] {
	return &ArrayList[T]{data: array}
}

func NewList[T any]() *ArrayList[T] {
	var array []T
	return &ArrayList[T]{data: array}
}

func (list *ArrayList[T]) Add(x T) { list.data = append([]T(list.data), x) }

func (list *ArrayList[T]) Remove(i int) {
	if i < len([]T(list.data)) {
		list.data = append([]T(list.data[:i]), list.data[i+1:]...)
	} else {
		panic(errors.New("array index out of bounds"))
	}
}

func (list *ArrayList[T]) Contains(elem T) bool {
	for _, item := range list.data {
		if reflect.ValueOf(item).Interface() == reflect.ValueOf(elem).Interface() {
			return true
		}
	}
	return false
}

func (list *ArrayList[T]) Count() int {
	return len([]T(list.data))
}

func (list *ArrayList[T]) ToSlice() []T {
	return list.data
}

func (list *ArrayList[T]) ToQueryable() Queryable[T] {
	return list.data
}
