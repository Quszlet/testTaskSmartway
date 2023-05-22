package service

import "github.com/Quszlet/testTaskSmartway/internal/models"

func (s *service) AddEmployee(employee models.Employee) (int, error) {
	return s.repository.AddEmployee(employee)
}

func (s *service) DeleteEmployee(id int) error {
	return s.repository.DeleteEmployee(id)
}

func (s *service) GetAllEmployeeCompany(company string) ([]models.Employee, error){
	return s.repository.GetAllEmployeeCompany(company)
}

func (s *service) GetAllEmployeeDepartament(company string, departmentId int) ([]models.Employee, error) {
	return s.repository.GetAllEmployeeDepartament(company, departmentId)
}

func (s *service) UpdateEmployee(id int, newData map[string]interface{}) error{
	return s.repository.UpdateEmployee(id, newData)
}