package routes

import (
	"projet/controllers"

	"github.com/gin-gonic/gin"
)

func PurchaseRoutes(r *gin.Engine) {
	r.GET("/api/orders/history", controllers.PurchasesList)
	r.GET("/api/orders/history/:idd", controllers.PurchaseList)
	r.POST("/api/orders/create", controllers.PurchaseAdd)
}
