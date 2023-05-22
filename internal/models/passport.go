package models

type Passport struct {
	Id int `json:"-" db:"id"`
	Type   string `json:"type" db:"type" binding:"required"`
	Number string `json:"number" db:"number" binding:"required"`
}