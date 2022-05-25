package glinq

type orderedSlice[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func (s orderedSlice[T]) Len() int           { return len([]T(s.s)) }
func (s orderedSlice[T]) Less(i, j int) bool { return s.cmp(s.s[i], s.s[j]) }
func (s orderedSlice[T]) Swap(i, j int)      { s.s[i], s.s[j] = s.s[j], s.s[i] }
