package main

import (
	"projet/config"
	"projet/routes"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true}))

	router.Static("/static", "./public")
	router.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	routes.AuthorRoutes(router)
	routes.BooksRoutes(router)
	routes.UserRoutes(router)
	routes.PurchaseRoutes(router)

	router.Run(":3001")
}
