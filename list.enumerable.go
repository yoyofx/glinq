package glinq

type ListEnumerable[T any] struct {
	list  []T
	index int
}

func (e *ListEnumerable[T]) Next() bool {
	e.index++
	if e.index < len(e.list) {
		return true
	}
	return false
}

func (e *ListEnumerable[T]) Reset() {
	e.index = 0
}

func (e *ListEnumerable[T]) Value() (T, bool) {
	if e.index >= 0 && e.index < len(e.list) {
		return e.list[e.index], true
	} else {
		var t T
		return t, false
	}
}
