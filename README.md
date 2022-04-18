# glinq
LINQ for Golang

examples:
```go
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
```