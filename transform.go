package glinq

// Reduce uses the summary function Queryable[T1] summary into one result of type T2.
func Reduce[T1, T2 any](from Queryable[T1], initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range from {
		r = f(r, v)
	}
	return r
}

// Map uses the mapping function to convert Queryable[T1] into Queryable[T2]
func Map[T1, T2 any](from Queryable[T1], f func(T1) T2) Queryable[T2] {
	r := make([]T2, from.Count())
	from.ForEach(func(i int, elem T1) {
		r[i] = f(elem)
	})
	return From(r)
}
