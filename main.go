package main

import (
	"fmt"
	"github.com/pushm0v/go-adventure/employee"
)

func main() {
	var es = employee.NewEmployeeService()
	es.NewEmployee(1, "Bherly", "Novrandy", "IT", 10000)

	fmt.Printf("%+v", es.GetEmployee(1))
}
