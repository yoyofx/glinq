package glinq

type Stack[T any] []T

func NewStackOf[T any](from []T) *Stack[T] {
	stack := Stack[T](from)
	return &stack
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(elem T) {
	*s = append(*s, elem)
}
func (s *Stack[T]) Pop() *T {
	elem := s.Top()
	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}
	return elem
}

func (s *Stack[T]) Top() *T {
	if len(*s) > 0 {
		return &(*s)[len(*s)-1]
	}
	return nil
}

func (s *Stack[T]) Count() int {
	return len(*s)
}

func (s *Stack[T]) ToSlice() []T {
	return *s
}

func (s *Stack[T]) ToQueryable() Queryable[T] {
	return s.ToSlice()
}

// GetEnumerator get enumerable object
func (s *Stack[T]) GetEnumerator() IEnumerator[T] {
	return &ListEnumerable[T]{list: *s, index: 0}
}
