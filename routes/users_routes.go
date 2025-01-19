package routes

import (
	"projet/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/api/users", controllers.UsersList)
	r.GET("/api/users/:idd", controllers.UserList)
	r.POST("/api/users/connect", controllers.UserConnect)
	r.POST("/api/users", controllers.UserAdd)
	r.PUT("/api/users/:idd", controllers.AdminModify)
	r.PUT("/api/users", controllers.UserModify)
	r.DELETE("/api/users/:idd", controllers.AdminDelete)
	r.DELETE("/api/users", controllers.UserDelete)
}
