package handler

import (
	"net/http"

	serv "github.com/Shin0kari/go_max"
	"github.com/gin-gonic/gin"
)

// аунтентификация
func (h *Handler) signUp(c *gin.Context) {
	// в этой структуре записываем данные о пользователях
	var input serv.User

	// c - gin context; c.BindJSON - принимает ссылку на объект, чтобы распарсить JSON
	if err := c.BindJSON(&input); err != nil {
		// функция для создания ответа с ошибкой
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// в обработчик передаём структуру пользователя`
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// регистрация
func (h *Handler) signIn(c *gin.Context) {
	// в этой структуре записываем данные о пользователях
	var input signInInput

	// c - gin context; c.BindJSON - принимает ссылку на объект, чтобы распарсить JSON
	if err := c.BindJSON(&input); err != nil {
		// функция для создания ответа с ошибкой
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// в обработчик передаём структуру пользователя`
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
