package routes

import (
	"projet/controllers"

	"github.com/gin-gonic/gin"
)

func BooksRoutes(r *gin.Engine) {
	r.GET("/api/books", controllers.BooksList)
	r.POST("/api/books/:idd", controllers.BookList)
	r.POST("/api/books", controllers.BookAdd)
	r.PUT("/api/books/:idd", controllers.BookModify)
	r.DELETE("/api/books/:idd", controllers.BookDelete)
}
