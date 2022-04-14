package glinq

import "sort"

type orderedSlice[T Number] []T

func (s orderedSlice[T]) Len() int           { return len([]T(s)) }
func (s orderedSlice[T]) Less(i, j int) bool { return s[i] < s[j] }
func (s orderedSlice[T]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type orderedDescSlice[T Number] []T

func (s orderedDescSlice[T]) Len() int           { return len([]T(s)) }
func (s orderedDescSlice[T]) Less(i, j int) bool { return s[i] > s[j] }
func (s orderedDescSlice[T]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func SortAsc[T Number](s []T) []T {
	sort.Sort(orderedSlice[T](s))
	return s
}

func SortDesc[T Number](s []T) []T {
	sort.Sort(orderedDescSlice[T](s))
	return s
}
