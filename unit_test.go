package glinq

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSort(t *testing.T) {
	s1 := From([]int{3, 5, 2})
	s1.Sort(func(a, b int) bool {
		return a > b
	})
	//s2 := From([]int{3, 5, 2})
	assert.Equal(t, s1.ToSlice(), []int{5, 3, 2})
	s1.Sort(func(a, b int) bool {
		return a < b
	})
	assert.Equal(t, s1.ToSlice(), []int{2, 3, 5})
}

func TestLinq(t *testing.T) {
	query := From([]int{0, 1, 2, 3, 4})
	one, _ := query.First()
	assert.Equal(t, one, 0)
	assert.Equal(t, query.CountIf(func(item int) bool { return item%2 == 0 }), 3)

	oQuery := From([]int{2, 4})
	assert.Equal(t, oQuery.All(func(item int) bool { return item%2 == 0 }), true)

	jQuery := From([]int{1, 2, 4})
	assert.Equal(t, jQuery.Any(func(item int) bool { return item%2 != 0 }), true)

	oSlice := From([]int{1, 2, 4}).Where(func(item int) bool { return item%2 == 0 })
	assert.Equal(t, oSlice.ToSlice(), []int{2, 4})

	qtFrom := From([]int{1, 2, 4})
	sum := Reduce(qtFrom, 0, func(i, j int) int { return i + j })
	assert.Equal(t, sum, 7)

	mapQuery := Map(qtFrom, func(item int) string {
		return strconv.Itoa(item)
	})
	assert.Equal(t, mapQuery.ToSlice(), []string{"1", "2", "4"})

}

func TestArrayList(t *testing.T) {
	list1 := NewListOf([]int{1, 2, 4})
	list1.Push(5)

	assert.Equal(t, list1.ToSlice(), []int{1, 2, 4, 5})
	list1.RemoveAt(3)
	// now list items of { 1, 2, 4}
	assert.Equal(t, list1.ToSlice(), []int{1, 2, 4})
	assert.Equal(t, list1.Contains(2), true)
	assert.Equal(t, list1.Contains(5), false)

	list1.RemoveAt(0)
	// now list items of { 2, 4}
	qlist := list1.ToQueryable()
	DoWhile(qlist.GetEnumerator(), func(v int) bool {
		if v == 2 {
			assert.Equal(t, v, 2)
		} else {
			assert.Equal(t, v, 4)
		}
		return true
	})
	assert.Equal(t, qlist.All(func(item int) bool { return item%2 == 0 }), true)

}

func TestListEnumerable(t *testing.T) {
	list1 := NewListOf([]int{1, 2, 4})
	it := list1.GetEnumerator()
	DoWhile(it, func(v int) bool {
		if v == 4 { // v==4 break
			return false
		}
		fmt.Println(v) // { 1, 2 }
		return true
	})
}

func TestICollectionAndEnumerator(t *testing.T) {
	var collection ICollection[int] = NewListOf([]int{1, 2, 3})
	DoWhile(collection.GetEnumerator(), func(v int) bool {
		fmt.Println(v)
		return true
	})

	var collection2 ICollection[int] = NewLinkedListOf([]int{2, 3, 4})
	DoWhile(collection2.GetEnumerator(), func(v int) bool {
		fmt.Println(v)
		return true
	})
}

func TestLinkedList(t *testing.T) {
	linkedList := NewLinkedListOf([]int{2, 3, 4})
	linkedList.Push(5)
	assert.Equal(t, linkedList.ToSlice(), []int{2, 3, 4, 5})
}
