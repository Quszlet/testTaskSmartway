package repository

import (
	"errors"
	"fmt"
	"github.com/Quszlet/testTaskSmartway/internal/models"
)

func (repo *Repository) AddEmployee(employee models.Employee) (int, error) {
	var (
		idEmployee   int
		idPassport   int
		idDepartment int
	)

	query := fmt.Sprintf("SELECT departments.id FROM %s JOIN %s ON companies.id = $1 WHERE departments.companyId = $1 AND departments.name = $2 AND departments.phone = $3", departmentsTable, companiesTable)

	row := repo.database.QueryRow(query, employee.CompanyId, employee.Department.Name, employee.Department.Phone)
	if err := row.Scan(&idDepartment); err != nil {
		return 0, errors.New("there is no department with that name or number in the company or there is no such company")
	}

	query = fmt.Sprintf("INSERT INTO %s (name, surname, phone, companyId, departmentId) VALUES ($1, $2, $3, $4, $5) RETURNING id", employeesTable)

	row = repo.database.QueryRow(query, employee.Name, employee.Surname, employee.Phone, employee.CompanyId, idDepartment)
	if err := row.Scan(&idEmployee); err != nil {

		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (type, number, employeeId) VALUES ($1, $2, $3) RETURNING id", passportsTable)

	row = repo.database.QueryRow(query, employee.Passport.Type, employee.Passport.Number, idEmployee)
	if err := row.Scan(&idPassport); err != nil {
		return 0, err
	}

	return idEmployee, nil
}

func (repo *Repository) DeleteEmployee(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE employees.id = $1", employeesTable)
	_, err := repo.database.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetAllEmployeeCompany(company string) ([]models.Employee, error) {
	var employees []models.Employee

	query := `SELECT employees.name, employees.surname, employees.phone, employees.companyId, 
			passports.type AS passport_type, passports.number AS passport_number, departments.name AS department_name, 
			departments.phone AS department_phone 
			FROM employees JOIN passports ON employees.id = passports.employeeId 
			JOIN departments ON employees.departmentId = departments.id 
			JOIN companies ON companies.id = employees.companyId WHERE companies.name = $1`
	rows, err := repo.database.Query(query, company)

	if err != nil {
		return nil, err
	}

	var (
		name             string
		surname          string
		phone            string
		company_id       int
		passport_type    string
		passport_number  string
		department_name  string
		department_phone string
	)


	for rows.Next() {

		err = rows.Scan(&name, &surname, &phone, &company_id, &passport_type, &passport_number, &department_name, &department_phone)

		if err != nil {
			return nil, err
		}

		passport := models.Passport{
			Type:   passport_type,
			Number: passport_number,
		}

		department := models.Department{
			Name:  department_name,
			Phone: department_phone,
		}

		employee := models.Employee{
			Name:       name,
			Surname:    surname,
			Phone:      phone,
			CompanyId:  company_id,
			Passport:   passport,
			Department: department,
		}

		employees = append(employees, employee)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (repo *Repository) GetAllEmployeeDepartament(company string, departmentId int) ([]models.Employee, error) {
	var employees []models.Employee

	query := `SELECT employees.name, employees.surname, employees.phone, employees.companyId, 
			passports.type AS passport_type, passports.number AS passport_number, departments.name AS department_name, 
			departments.phone AS department_phone 
			FROM employees JOIN passports ON employees.id = passports.employeeId 
			JOIN departments ON employees.departmentId = departments.id 
			JOIN companies ON companies.id = employees.companyId WHERE departments.id = $1 AND companies.name = $2`
	rows, err := repo.database.Query(query, departmentId, company)

	if err != nil {
		return nil, err
	}

	var (
		name             string
		surname          string
		phone            string
		company_id       int
		passport_type    string
		passport_number  string
		department_name  string
		department_phone string
	)

	for rows.Next() {

		err = rows.Scan(&name, &surname, &phone, &company_id, &passport_type, &passport_number, &department_name, &department_phone)

		if err != nil {
			return nil, err
		}

		passport := models.Passport{
			Type:   passport_type,
			Number: passport_number,
		}

		department := models.Department{
			Name:  department_name,
			Phone: department_phone,
		}

		employee := models.Employee{
			Name:       name,
			Surname:    surname,
			Phone:      phone,
			CompanyId:  company_id,
			Passport:   passport,
			Department: department,
		}

		employees = append(employees, employee)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (repo *Repository) UpdateEmployee(id int, newData map[string]interface{}) error {
	for key, value := range newData {
		var err error
		switch key {
		case "name", "surname", "phone":
			query := fmt.Sprintf("UPDATE %s SET %s = '%s' WHERE id = $1", employeesTable, key, value)
			_, err = repo.database.Exec(query, id)
		case "companyId":
			query := fmt.Sprintf("UPDATE %s SET %s = %f WHERE id = $1", employeesTable, key, value)
			_, err = repo.database.Exec(query, id)
		case "passport":
			for keyPassport, valuePassport := range value.(map[string]interface{}) {
				query := fmt.Sprintf("UPDATE %s SET %s = '%s' WHERE %s.employeeid = $1", passportsTable,
						 keyPassport, valuePassport, passportsTable)
					_, err = repo.database.Exec(query, id)
			}
		case "department":
			tx, errCommit := repo.database.Begin()
			if errCommit != nil {
				return err
			}

			for keyDepartment, valueDepartment := range value.(map[string]interface{}) {
				switch keyDepartment {
				case "name":
					query := `UPDATE employees SET departmentid = departments.id FROM departments, companies 
								WHERE departments.name = $1
								AND employees.id = $2 
								AND employees.companyid = departments.companyid`

					result, err := tx.Exec(query, valueDepartment, id)
					if err != nil {
						return err
					}

					rowsAffected, err := result.RowsAffected()
					if err != nil {
						tx.Rollback()
						return err
					}
				
					if rowsAffected == 0 {
						tx.Rollback()
						return errors.New("update failed, there is no department with that name in the company where the employee belongs")
					}

				case "phone":
					query := `UPDATE employees SET departmentid = departments.id FROM departments, companies 
								WHERE departments.phone = $1
								AND employees.id = $2 
								AND employees.companyid = departments.companyid`

					result, err := tx.Exec(query, valueDepartment, id)
					if err != nil {
						return err
					}

					rowsAffected, err := result.RowsAffected()
					if err != nil {
						tx.Rollback()
						return err
					}
				
					if rowsAffected == 0 {
						tx.Rollback()
						return errors.New("update failed, there is no department with that number in the company where the employee belongs")
					}
				}
				
			}
			errCommit = tx.Commit()
			if errCommit != nil {
				tx.Rollback()
				return errCommit
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
