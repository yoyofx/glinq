package glinq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	s1 := []int{3, 5, 2}
	s2 := []int{3, 5, 2}
	assert.Equal(t, SortDesc(s1), []int{5, 3, 2})
	assert.Equal(t, SortAsc(s2), []int{2, 3, 5})
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
}
