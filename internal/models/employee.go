package models

type Employee struct {
	Id         int        `json:"-"`
	Name       string     `json:"name" binding:"required"`
	Surname    string     `json:"surname" binding:"required"`
	Phone      string     `json:"phone" binding:"required"`
	CompanyId  int        `json:"companyid" binding:"required"`
	Passport   Passport   `json:"passport" binding:"required"`
	Department Department `json:"department" binding:"required"`
}
