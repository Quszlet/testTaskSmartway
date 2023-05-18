package service

type EmployeeService interface {
	AddEmployee() 
	DeleteEmployee(id int)
	GetAllEmployeeCompany(company string)
	GetAllEmployeeDepartament(departament string)
	UpdateEmployee(id int, newData map[string]string)
}

type service struct {
	EmployeeService
}

func NewService() *service {
	return &service{
		
	}
}