package router

import (
	"github.com/Komei1009/todoList/controller"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/signup",controller.CreateUserController)
	api.POST("/signin",controller.LoginController)

	api.POST("/addTodo", controller.TodoAdd)
	api.POST("/controllerTodo", controller.TodoControll)
	api.POST("/todoDisplay", controller.TodoDisplay)
}
