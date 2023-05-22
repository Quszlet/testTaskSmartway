package handler

import (
	"github.com/Quszlet/testTaskSmartway/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.EmployeeService
}

func NewHandeler(service service.EmployeeService) *Handler{
	return &Handler{
		services: service,
	}
}

func (h *Handler) SetupRoutes(gin *gin.Engine) {
	api := gin.Group("/api")
	{
		employees := api.Group("/employee")
		{
			employees.POST("/", h.AddEmployee)
			employees.DELETE("/:id", h.DeleteEmployee)
			employees.GET("/company/:company_name", h.GetAllEmployeeCompany)
			employees.GET("/company/:company_name/department/:department_id", h.GetAllEmployeeDepartament)
			employees.PUT("/:id", h.UpdateEmployee)
		}
	}


}