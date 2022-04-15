package glinq

import "sort"

// Queryable query collection
type Queryable[T any] []T

func From[T any](items []T) Queryable[T] {
	return items
}

// ToSlice Queryable to slice
func (query Queryable[T]) ToSlice() []T {
	return query
}

// Where filters a collection of values based on a predicate.
func (query Queryable[T]) Where(predicate func(T) bool) Queryable[T] {
	var result []T
	for _, elem := range query {
		if predicate(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// All determines whether all elements of a collection satisfy a condition.
func (query Queryable[T]) All(predicate func(T) bool) bool {
	for _, elem := range query {
		if !predicate(elem) {
			return false
		}
	}
	return true
}

// Any determines whether any element of a collection satisfies a condition.
func (query Queryable[T]) Any(predicate func(T) bool) bool {
	for _, elem := range query {
		if predicate(elem) {
			return true
		}
	}
	return false
}

// ForEach performs the specified action on each element of a collection.
func (query Queryable[T]) ForEach(action func(int, T)) {
	for index, elem := range query {
		action(index, elem)
	}
}

func (query Queryable[T]) ToChannel(result chan<- T) {
	query.ForEach(func(index int, elem T) {
		result <- elem
	})
}

// Count this queryable collection's count.
func (query Queryable[T]) Count() int {
	return len([]T(query))
}

// CountIf count of collection satisfy a condition.
func (query Queryable[T]) CountIf(predicate func(T) bool) int {
	return query.Where(predicate).Count()
}

// First return the first element of a collection.
func (query Queryable[T]) First() (T, error) {
	var item T
	slice := query.ToSlice()
	if len([]T(slice)) > 0 {
		item = slice[0]
	}
	return item, ErrorCannotFound
}

// FirstIf return first element of collection satisfy a condition.
func (query Queryable[T]) FirstIf(predicate func(T) bool) (T, error) {
	return query.Where(predicate).First()
}

// Sort return the sort element of a collection by compare function.
func (query Queryable[T]) Sort(cmp func(T, T) bool) {
	sort.Sort(orderedSlice[T]{query, cmp})
}
