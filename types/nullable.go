package types

import "fmt"

type Nullable[T any] struct {
	hasValue bool
	value    T
}

func NullableOf[T](v *T) Nullable[T] {
	return Nullable[T]{value: *v, hasValue: v != nil}
}

func (nullable Nullable[T]) GetValue() T {
	return nullable.value
}

func (nullable Nullable[T]) GetPointer() *T {
	return &nullable.value
}

func (nullable Nullable[T]) HasValue() bool {
	return nullable.hasValue
}

func (nullable Nullable[T]) String() string {
	stringer := nullable.value.(fmt.Stringer)
	if stringer != nil {
		return stringer.String()
	}
	return ""
}
