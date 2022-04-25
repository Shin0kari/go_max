package handler

import (
	"github.com/Shin0kari/go_max/package/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// инициализирует endpoints(как я понял, функции)
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// объявляем методы, пронумеровав их маршрутами
	auth := router.Group("/auth")
	{
		// авторизация
		auth.POST("/sign-up", h.signUp)
		// регистрация
		auth.POST("/sign-in", h.signIn)
	}

	//для работы endpoints с дисками и их задачами
	api := router.Group("/api", h.userIndentity)
	{
		// группа для работы со списками
		lists := api.Group("/lists")
		{
			// создание получение редактирование удаление списков
			lists.POST("/", h.createList)
			lists.GET("/", h.getALLLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			// группа для задач списка
			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
