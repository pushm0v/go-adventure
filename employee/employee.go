package employee

type Employee struct {
	ID int
	FirstName string
	LastName string
	Department string
	Salary int
	LeavesTaken int
}

type employeeService struct {
	employees []Employee
}

type EmployeeInterface interface {
	NewEmployee(id int, firstName, lastName, department string, salary int)
	TakeLeave(id int, amount int)
	LeavesCount(id int) int
	UpdateEmployee(id int, employee Employee)
	DeductSalary(id int, amount int)
	IncreaseSalary(id int, amount int)
	GetEmployee(id int) Employee
}

func NewEmployeeService() EmployeeInterface {
	return &employeeService{}
}

func (es *employeeService) NewEmployee(id int, firstName, lastName, department string, salary int) {
	var e = Employee{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Department:  department,
		Salary:      salary,
		LeavesTaken: 0,
	}

	es.employees = append(es.employees, e)
}

func (es *employeeService) TakeLeave(id int, amount int) {
	var e = es.findEmployee(id)
	e.LeavesTaken = e.LeavesTaken + amount
}

func (es *employeeService) LeavesCount(id int) int {
	var e = es.findEmployee(id)
	return e.LeavesTaken
}

func (es *employeeService) UpdateEmployee(id int, employee Employee) {
	var e = es.findEmployee(id)
	e.FirstName = employee.FirstName
	e.LastName = employee.LastName
	e.Department = employee.Department
	e.Salary = employee.Salary
}

func (es *employeeService) DeductSalary(id int, amount int) {
	var e = es.findEmployee(id)
	e.Salary = e.Salary - amount
}

func (es *employeeService) IncreaseSalary(id int, amount int) {
	var e = es.findEmployee(id)
	e.Salary = e.Salary + amount
}

func (es *employeeService) GetEmployee(id int) Employee {
	return *es.findEmployee(id)
}

func (es *employeeService) findEmployee(id int) *Employee {
	var employee *Employee
	for k, e := range es.employees {
		if e.ID == id {
			employee = &es.employees[k]
			break
		}
	}

	return employee
}
