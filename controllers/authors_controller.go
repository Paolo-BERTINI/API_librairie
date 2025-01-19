package controllers

import (
	"log"
	"net/http"
	"projet/config"
	"projet/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func AuthorList(c *gin.Context) {
	id := c.Param("idd")
	rows, err := config.DB.QueryContext(c, "SELECT name, birth_date, description, id FROM authors WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	defer rows.Close()
	var authors []models.Author_list
	for rows.Next() {
		var author models.Author_list
		if err := rows.Scan(&author.Name, &author.Birth_date, &author.Description, &author.Id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		authors = append(authors, author)
	}
	c.JSON(http.StatusOK, gin.H{"Authors": authors})
}

func AuthorsList(c *gin.Context) {
	rows, err := config.DB.QueryContext(c, "SELECT name FROM authors")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	defer rows.Close()
	var authors []models.Author_name
	for rows.Next() {
		var author models.Author_name
		if err := rows.Scan(&author.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		authors = append(authors, author)
	}
	c.JSON(http.StatusOK, gin.H{"Authors": authors})
}

func AuthorAdd(c *gin.Context) {
	if Id_user == "1" {
		var NewAuthor map[string]string
		if err := c.ShouldBindJSON(&NewAuthor); err != nil {
			c.JSON(400, gin.H{"Error": err.Error()})
			return
		}
		var name, birth_date, description string
		if val, ok := NewAuthor["name"]; ok {
			name = val
		}
		if val, ok := NewAuthor["birth_date"]; ok {
			birth_date = val
		}
		if val, ok := NewAuthor["description"]; ok {
			description = val
		}
		result, err := config.DB.ExecContext(c,
			"INSERT INTO authors (name, birth_date, description) VALUES(?, ?, ?)", name, birth_date, description)
		if err != nil {
			log.Fatal("Erreur lors de la récupération des auteurs : ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Erreur lors de la récupération du nombre de lignes affectées : ", err)
			return
		}
		log.Printf("Nombre de lignes affectées : %d", rowsAffected)
		c.JSON(http.StatusOK, gin.H{
			"Message":                "Auteur créé avec succès",
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

func AuthorModify(c *gin.Context) {
	if Id_user == "1" {
		id := c.Param("idd")
		var NewAuthor map[string]string
		if err := c.ShouldBindJSON(&NewAuthor); err != nil {
			c.JSON(400, gin.H{"Error": err.Error()})
			return
		}
		var name, birth_date, description string
		if val, ok := NewAuthor["name"]; ok {
			name = val
		}
		if val, ok := NewAuthor["birth_date"]; ok {
			birth_date = val
		}
		if val, ok := NewAuthor["description"]; ok {
			description = val
		}
		result, err := config.DB.ExecContext(c,
			"UPDATE authors SET name = ?, birth_date = ?, description = ? WHERE id = ?", name, birth_date, description, id)
		if err != nil {
			log.Fatal("Erreur lors de la récupération des auteurs : ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Erreur lors de la récupération du nombre de lignes affectées : ", err)
			return
		}
		log.Printf("Nombre de lignes affectées : %d", rowsAffected)
		c.JSON(http.StatusOK, gin.H{
			"Message":                   "Auteur modifiée avec succès",
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

func AuthorDelete(c *gin.Context) {
	if Id_user == "1" {
		id := c.Param("idd")
		result, err := config.DB.ExecContext(c,
			"DELETE FROM authors WHERE id = ?", id)
		if err != nil {
			log.Fatal("Erreur lors de la récupération des auteurs : ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Erreur lors de la récupération du nombre de lignes affectées : ", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message":                    "Auteur supprimée avec succès",
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
