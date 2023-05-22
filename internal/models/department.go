package models

type Department struct {
	Id int `json:"-" db:"id"`
	Name  string `json:"name" db:"department_name" binding:"required"`
	Phone string `json:"phone" db:"department_phone" binding:"required"`
}