package handler

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, message string, statusCode int) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}