package models

type Department struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}
