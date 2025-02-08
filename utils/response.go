package utils

import (
	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func SendErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  "error",
		"message": message,
	})
}
