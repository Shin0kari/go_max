package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//функция для обработки ошибок
type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}
