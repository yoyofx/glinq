package glinq

import (
	"errors"
	"reflect"
)

type ICollection[T any] interface {
	Push(T)
	RemoveAt(int)
	Remove(T)
	Count() int
	Contains(T) bool
	IndexOf(T) int
	GetEnumerator() IEnumerator[T]
}

type IEnumerator[T any] interface {
	Next() bool
	Reset()
	Value() (T, bool)
}

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type Ordered interface {
	Number | string
}

var (
	ErrorCannotFound      = errors.New("can not find element")
	ErrorIndexoutOfBounds = errors.New("index out of bounds")
)

func Equal[T1, T2 any](a T1, b T2) bool {
	return reflect.ValueOf(a).Interface() == reflect.ValueOf(b).Interface()
}

func Sum[T Number](query Queryable[T]) T {
	var sum T
	for _, elem := range query {
		sum += elem
	}
	return sum
}
