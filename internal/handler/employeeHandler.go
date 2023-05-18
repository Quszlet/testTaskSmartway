package handler

import (
	"fmt"
	"net/http"

	"github.com/Quszlet/testTaskSmartway/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddEmployee(c *gin.Context) {
	var request models.Employee

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, "invalid json", http.StatusBadRequest)
		return
	}

	fmt.Println("123")
}

func (h *Handler) DeleteEmployee(c *gin.Context) {

}

func (h *Handler) GetAllEmployeeCompany(c *gin.Context) {

}

func (h *Handler) GetAllEmployeeDepartament(c *gin.Context) {

}

func (h *Handler) UpdateEmployee(c *gin.Context) {

}