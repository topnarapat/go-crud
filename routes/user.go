package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/topnarapat/go-crud/controllers"
)

func InitUserRoutes(r *gin.Engine) {
	r.GET("/users", controllers.GetAllUsers)

	r.POST("/register", controllers.Register)
}
