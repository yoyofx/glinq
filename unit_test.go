package glinq

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	g "github.com/yoyofx/glinq/generic"
	"strconv"
	"testing"
)

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   float64
}

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
	qList := list1.ToQueryable()
	DoWhile(qList.GetEnumerator(), func(v int) bool {
		if v == 2 {
			assert.Equal(t, v, 2)
		} else {
			assert.Equal(t, v, 4)
		}
		return true
	})
	assert.Equal(t, qList.All(func(item int) bool { return item%2 == 0 }), true)

	justQuery := Just(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	assert.Equal(t, justQuery.Take(2).ToSlice(), []int{1, 2})

	assert.Equal(t, Just(1, 2, 3, 4, 5).Skip(2).ToSlice(), []int{3, 4, 5})

	assert.Equal(t, Range(0, 3).ToSlice(), []int{0, 1, 2, 3})

	queryDistinct := Just(1, 2, 2, 3, 4, 4, 5).Distinct()
	queryDistinct.Sort(func(a, b int) bool { return a < b })
	assert.Equal(t, queryDistinct.ToSlice(), []int{1, 2, 3, 4, 5})
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
	linkedList.Remove(3)
	assert.Equal(t, linkedList.ToSlice(), []int{2, 4, 5})
	assert.Equal(t, linkedList.IndexOf(4), 1)
	assert.Equal(t, linkedList.Count(), 3)
	assert.Equal(t, linkedList.Contains(6), false)
	assert.Equal(t, linkedList.IsEmpty(), false)
}

func TestStack(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("hello")
	stack.Push("world")
	assert.Equal(t, *stack.Pop(), "world")
	assert.Equal(t, *stack.Pop(), "hello")

	stack1 := NewStackOf([]string{"hello", "world", "everyone"})
	assert.Equal(t, *stack1.Pop(), "everyone")
	assert.Equal(t, stack1.ToSlice(), []string{"hello", "world"})
}

var employees = NewListOf([]Employee{
	{"Hao", 44, 0, 8000.5},
	{"Bob", 36, 10, 5000.5},
	{"Alice", 23, 5, 9000.0},
	{"Jack", 26, 0, 4000.0},
	{"Tom", 48, 9, 7500.75},
	{"Marry", 29, 0, 6000.0},
	{"Mike", 38, 8, 4000.3},
})

func TestEmployeeList(t *testing.T) {
	// 添加 Max 员工
	employees.Push(Employee{"Max", 26, 0, 4000.0})
	// 所有人的薪水
	totalPaySalary := Reduce(employees.ToQueryable(), 0.0, func(result float64, employee Employee) float64 {
		return result + employee.Salary
	})
	assert.Equal(t, totalPaySalary, 43502.05+4000.0)

	//统计年龄大于40岁的员工数
	orderCount := employees.ToQueryable().Where(func(employee Employee) bool {
		return employee.Age > 35
	}).Count()
	assert.Equal(t, orderCount, 4)

	//统计薪水超过 6000元的员工数
	moreThan6000Count := employees.ToQueryable().Where(func(employee Employee) bool {
		return employee.Salary >= 6000
	}).Count()
	assert.Equal(t, moreThan6000Count, 4)

	//统计年龄小于30岁员工要支付的所有薪水 [ Max: 4000.0, Alice: 9000.0 ,Jack: 4000.0, Marry: 6000.0 ] = 23000.0
	employeeQuery := employees.ToQueryable()
	// Reduce are equal Sum of Map to Queryable
	youngerTotalPaySalary := Reduce(Map(employeeQuery.Where(func(employee Employee) bool {
		return employee.Age < 30
	}), func(e Employee) float64 {
		return e.Salary
	}), 0.0, func(result float64, item float64) float64 {
		return result + item
	})
	assert.Equal(t, youngerTotalPaySalary, 23000.0)
}

func TestBTree(t *testing.T) {
	tree := NewBTree[int, string](g.Less[int])

	tree.Put(42, "foo")
	tree.Put(-10, "bar")
	tree.Put(0, "baz")

	tree.Each(func(key int, val string) {
		fmt.Println(key, val)
	})

}
