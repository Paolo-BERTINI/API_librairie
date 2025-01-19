package controllers

import (
	"net/http"
	"projet/config"
	"projet/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func panier(c *gin.Context) {

}

func PurchasesList(c *gin.Context) {
	if Id_user == "" {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Veuillez vous connecter"})
	} else {
		rows, err := config.DB.QueryContext(c, "SELECT * FROM purchase WHERE user = ?", Id_user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		defer rows.Close()
		var purchases []models.Purchase
		for rows.Next() {
			var purchase models.Purchase
			if err := rows.Scan(&purchase.Id, &purchase.User, &purchase.Quantity, &purchase.Total_price, &purchase.Payment_timestamp, &purchase.Book); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
			purchases = append(purchases, purchase)
		}
		c.JSON(http.StatusOK, gin.H{"Purchase": purchases})
	}
}

func PurchaseList(c *gin.Context) {
	if Id_user == "1" {
		id := c.Param("idd")
		rows, err := config.DB.QueryContext(c, "SELECT * FROM purchase WHERE user = ?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		defer rows.Close()
		var purchases []models.Purchase
		for rows.Next() {
			var purchase models.Purchase
			if err := rows.Scan(&purchase.Id, &purchase.User, &purchase.Quantity, &purchase.Total_price, &purchase.Payment_timestamp, &purchase.Book); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
			purchases = append(purchases, purchase)
		}
		c.JSON(http.StatusOK, gin.H{"Purchase": purchases})
	} else if Id_user > "1" {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Permission refusée"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Veuillez vous connecter"})
	}
}

func PurchaseAdd(c *gin.Context) {
	if Id_user > "0" {
		currentTime := time.Now()
		payment_timestamp := currentTime.Format("2006-01-02")
		var NewPurchase map[string]string
		if err := c.ShouldBindJSON(&NewPurchase); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var quantityStr, book string
		if val, ok := NewPurchase["quantity"]; ok {
			quantityStr = val
		}
		if val, ok := NewPurchase["book"]; ok {
			book = val
		}
		var price int
		rows, err := config.DB.QueryContext(c, "SELECT price FROM books WHERE id = ?", book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&price); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
		}
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			c.JSON(400, gin.H{"Error": "Quantité invalide, doit être un nombre entier"})
			return
		}
		var quantite int
		rows2, err := config.DB.QueryContext(c, "SELECT stock FROM books WHERE id = ?", book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		defer rows2.Close()
		for rows2.Next() {
			if err := rows2.Scan(&quantite); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
		}
		if quantite >= quantity {
			config.DB.ExecContext(c,
				"INSERT INTO purchase (user, quantity, total_price, payment_timestamp, book) VALUES(?, ?, ?, ?, ?)", Id_user, quantity, price*quantity, payment_timestamp, book)
			config.DB.ExecContext(c, "UPDATE books SET stock = stock - ? WHERE id = ?", quantity, book)
			c.JSON(http.StatusOK, gin.H{
				"Message": "Achat effectué",
			})
		}
		if quantite < quantity {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Quantité disponible insuffisante"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Veuillez vous connecter",
		})
	}
}
