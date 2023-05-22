package service

import (
	"github.com/Quszlet/testTaskSmartway/internal/models"
	"github.com/Quszlet/testTaskSmartway/internal/repository"
)

type EmployeeService interface {
	AddEmployee(employee models.Employee) (int, error)
	DeleteEmployee(id int) error
	GetAllEmployeeCompany(company string) ([]models.Employee, error)
	GetAllEmployeeDepartament(company string, departmentId int) ([]models.Employee, error)
	UpdateEmployee(id int, newData map[string]interface{}) error
}

type service struct {
	repository *repository.Repository
}

func NewService(repo *repository.Repository) *service {
	return &service{
		repository: repo,
	}
}