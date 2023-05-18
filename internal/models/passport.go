package models

type Passport struct {
	Type   string `json:"type" binding:"required"`
	Number string `json:"number" binding:"required"`
}
