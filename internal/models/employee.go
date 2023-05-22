package models

type Employee struct {
	Id         int        `json:"-" db:"id"`
	Name       string     `json:"name" db:"name" binding:"required"`
	Surname    string     `json:"surname" db:"surname" binding:"required"`
	Phone      string     `json:"phone" db:"phone" binding:"required"`
	CompanyId  int        `json:"companyId" db:"company_id" binding:"required"`
	Passport   Passport   `json:"passport" binding:"required"`
	Department Department `json:"department" binding:"required"`
}
