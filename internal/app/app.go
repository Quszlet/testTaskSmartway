package app

import (
	"github.com/Quszlet/testTaskSmartway/internal/handler"
	"github.com/Quszlet/testTaskSmartway/internal/service"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	gin := gin.Default()

	service := service.NewService()

	handler := handler.NewHandeler()

	handler.SetupRoutes(gin)

	gin.Run()
}