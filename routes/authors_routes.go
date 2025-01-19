package routes

import (
	"projet/controllers"

	"github.com/gin-gonic/gin"
)

func AuthorRoutes(r *gin.Engine) {
	r.GET("/api/authors", controllers.AuthorsList)
	r.POST("/api/authors/:idd", controllers.AuthorList)
	r.POST("/api/authors", controllers.AuthorAdd)
	r.PUT("/api/authors/:idd", controllers.AuthorModify)
	r.DELETE("/api/authors/:idd", controllers.AuthorDelete)
}
