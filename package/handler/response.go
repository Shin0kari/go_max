package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//функция для обработки ошибок
type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
