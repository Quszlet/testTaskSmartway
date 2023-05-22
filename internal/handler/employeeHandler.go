package handler

import (
	"errors"
	"net/http"
	"strconv"
	"github.com/Quszlet/testTaskSmartway/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddEmployee(c *gin.Context) {
	var request models.Employee

	if err := c.BindJSON(&request); err != nil {
		ErrorResponse(c, "invalid json", http.StatusBadRequest)
		return
	}

	id, err := h.services.AddEmployee(request)
	if err != nil {
		ErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	employeeId, err := strconv.Atoi(id)
	if err != nil {
		ErrorResponse(c, errors.New("employee id is not valid").Error(), http.StatusBadRequest)
		return
	}
	err = h.services.DeleteEmployee(employeeId)
	if err != nil {
		ErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
	})
}

func (h *Handler) GetAllEmployeeCompany(c *gin.Context) {
	company := c.Param("company_name")

	employees, err := h.services.GetAllEmployeeCompany(company)
	if err != nil {
		ErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Employees": employees,
	})
}

func (h *Handler) GetAllEmployeeDepartament(c *gin.Context) {
	company := c.Param("company_name")
	department := c.Param("department_id")
	departmentId, err := strconv.Atoi(department)
	if err != nil {
		ErrorResponse(c, errors.New("department id is not valid").Error(), http.StatusBadRequest)
		return
	}

	employees, err := h.services.GetAllEmployeeDepartament(company, departmentId)
	if err != nil {
		ErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Employees": employees,
	})
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	employeeId, err := strconv.Atoi(id)
	if err != nil {
		ErrorResponse(c, errors.New("employee id is not valid").Error(), http.StatusBadRequest)
		return
	}

	var newData map[string]interface{}
	if err := c.BindJSON(&newData); err != nil {
		ErrorResponse(c, errors.New("data is not valid").Error(), http.StatusBadRequest)
		return
	}
	
	err = h.services.UpdateEmployee(employeeId, newData)
	if err != nil {
		ErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
	})
}