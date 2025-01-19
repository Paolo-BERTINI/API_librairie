package controllers

import (
	"log"
	"net/http"
	"projet/config"
	"projet/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func BooksList(c *gin.Context) {
	if Id_user != "1" {
		rows, err := config.DB.QueryContext(c, "SELECT b.title, a.name, b.id FROM books b INNER JOIN authors a ON b.author=a.id WHERE stock>0")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		defer rows.Close()
		var books []models.BookResponse
		for rows.Next() {
			var book models.BookResponse
			if err := rows.Scan(&book.Title, &book.Author, &book.Id); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
			books = append(books, book)
		}
		c.JSON(http.StatusOK, gin.H{"Books": books})
	} else {
		rows, err := config.DB.QueryContext(c, "SELECT b.title, a.name, b.stock FROM books b INNER JOIN authors a ON b.author = a.id WHERE stock > 0")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		defer rows.Close()
		var books []models.BookResponseAdmin
		for rows.Next() {
			var book models.BookResponseAdmin
			if err := rows.Scan(&book.Title, &book.Author, &book.Stock); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
			books = append(books, book)
		}
		c.JSON(http.StatusOK, gin.H{"Books": books})
	}
}

func BookList(c *gin.Context) {
	if Id_user != "1" {
		id := c.Param("idd")
		rows, err := config.DB.QueryContext(c, "SELECT b.title, a.name, b.summary, b.price, b.stock, b.author FROM books b INNER JOIN authors a ON b.author=a.id WHERE b.id=?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		defer rows.Close()
		var books []models.BookResponses
		var stock int
		for rows.Next() {
			var book models.BookResponses
			if err := rows.Scan(&book.Title, &book.Author, &book.Summary, &book.Price, &stock, &book.Authorid); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
			books = append(books, book)
		}
		if stock > 0 {
			c.JSON(http.StatusOK, gin.H{"Books": books})
		}
		if stock < 1 {
			c.JSON(http.StatusOK, gin.H{"Books": "Plus en stock"})
		}
	} else {
		id := c.Param("idd")
		rows, err := config.DB.QueryContext(c, "SELECT b.title, a.name, b.summary, b.price, b.stock FROM books b INNER JOIN authors a ON b.author=a.id WHERE b.id=?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		defer rows.Close()
		var books []models.BookResponsesAdmin
		for rows.Next() {
			var book models.BookResponsesAdmin
			if err := rows.Scan(&book.Title, &book.Author, &book.Summary, &book.Price, &book.Stock); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
			books = append(books, book)
		}
		c.JSON(http.StatusOK, gin.H{"Books": books})
	}
}

func BookAdd(c *gin.Context) {
	if Id_user == "1" {
		var NewBook map[string]string
		if err := c.ShouldBindJSON(&NewBook); err != nil {
			c.JSON(400, gin.H{"Error": err.Error()})
			return
		}
		var title, author, stock, publication_date, price, summary string
		if val, ok := NewBook["title"]; ok {
			title = val
		}
		if val, ok := NewBook["authors"]; ok {
			author = val
		}
		if val, ok := NewBook["stock"]; ok {
			stock = val
		}
		if val, ok := NewBook["publication_date"]; ok {
			publication_date = val
		}
		if val, ok := NewBook["price"]; ok {
			price = val
		}
		if val, ok := NewBook["summary"]; ok {
			summary = val
		}
		result, err := config.DB.ExecContext(c,
			"INSERT INTO books (title, author, stock, publication_date, price, summary) VALUES(?, ?, ?, ?, ?, ?)", title, author, stock, publication_date, price, summary)
		if err != nil {
			log.Fatal("Erreur lors de la récupération des livres : ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Erreur lors de la récupération du nombre de lignes affectées : ", err)
			return
		}
		log.Printf("Nombre de lignes affectées : %d", rowsAffected)
		c.JSON(http.StatusOK, gin.H{
			"Message":                "Livre créé avec succès",
			"Nombre de lignes créés": rowsAffected,
		})
	} else if Id_user > "1" {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Permission refusée"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Veuillez vous connecter"})
	}
}

func BookModify(c *gin.Context) {
	if Id_user == "1" {
		var NewBook map[string]string
		id := c.Param("idd")
		if err := c.ShouldBindJSON(&NewBook); err != nil {
			c.JSON(400, gin.H{"Error": err.Error()})
			return
		}
		var title, author, stock, publication_date, price, summary string
		if val, ok := NewBook["title"]; ok {
			title = val
		}
		if val, ok := NewBook["author"]; ok {
			author = val
		}
		if val, ok := NewBook["stock"]; ok {
			stock = val
		}
		if val, ok := NewBook["publication_date"]; ok {
			publication_date = val
		}
		if val, ok := NewBook["price"]; ok {
			price = val
		}
		if val, ok := NewBook["summary"]; ok {
			summary = val
		}
		result, err := config.DB.ExecContext(c,
			"UPDATE books SET title = ?, author = ?, stock = ?, publication_date = ?, price = ?, summary = ? WHERE id = ?", title, author, stock, publication_date, price, summary, id)
		if err != nil {
			log.Fatal("Erreur lors de la récupération des livres : ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Erreur lors de la récupération du nombre de lignes affectées : ", err)
			return
		}
		log.Printf("Nombre de lignes affectées : %d", rowsAffected)
		c.JSON(http.StatusOK, gin.H{
			"Message":                   "Livre modifiée avec succès",
			"Nombre de lignes modifiés": rowsAffected,
		})
	} else if Id_user > "1" {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Permission refusée"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Veuillez vous connecter"})
	}
}

func BookDelete(c *gin.Context) {
	if Id_user == "1" {
		id := c.Param("idd")
		result, err := config.DB.ExecContext(c,
			"DELETE FROM books WHERE id = ?", id)
		if err != nil {
			log.Fatal("Erreur lors de la récupération des livres : ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Erreur lors de la récupération du nombre de lignes affectées : ", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message":                    "Livre supprimée avec succès",
			"Nombre de lignes supprimés": rowsAffected,
		})
	} else if Id_user > "1" {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Permission refusée"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Veuillez vous connecter"})
	}
}
