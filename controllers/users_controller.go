package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"projet/config"
	"projet/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func UsersList(c *gin.Context) {
	if Id_user > "1" {
		rows, err := config.DB.QueryContext(c, "SELECT first_name, last_name, mail FROM users WHERE id = ?", Id_user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		defer rows.Close()
		var users []models.User_info
		for rows.Next() {
			var user models.User_info
			if err := rows.Scan(&user.First_name, &user.Last_name, &user.Mail); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
			users = append(users, user)
		}
		c.JSON(http.StatusOK, gin.H{"users": users})
	} else if Id_user == "1" {
		rows, err := config.DB.QueryContext(c, "SELECT id, first_name, last_name, mail FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		defer rows.Close()
		var users []models.User_info_admin
		for rows.Next() {
			var user models.User_info_admin
			if err := rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.Mail); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
			users = append(users, user)
		}
		c.JSON(http.StatusOK, gin.H{"Users": users})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Veuillez vous connecter"})
	}
}

func UserList(c *gin.Context) {
	if Id_user == "1" {
		id := c.Param("idd")
		rows, err := config.DB.QueryContext(c, "SELECT id, first_name, last_name, mail FROM users WHERE id = ?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		defer rows.Close()
		var users []models.User
		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.Mail); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
			users = append(users, user)
		}
		c.JSON(http.StatusOK, gin.H{"Users": users})
	} else if Id_user > "1" {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Permission refusée"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Veuillez vous connecter"})
	}
}

func UserAdd(c *gin.Context) {
	var NewUser map[string]string
	if err := c.ShouldBindJSON(&NewUser); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	var first_name, last_name, mail, password string
	if val, ok := NewUser["first_name"]; ok {
		first_name = val
	}
	if val, ok := NewUser["last_name"]; ok {
		last_name = val
	}
	if val, ok := NewUser["mail"]; ok {
		mail = val
	}
	if val, ok := NewUser["password"]; ok {
		password = val
	}
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Erreur lors du hachage du mot de passe"})
		return
	}
	result, err := config.DB.ExecContext(c,
		"INSERT INTO users (first_name, last_name, mail, password) VALUES(?, ?, ?, ?)", first_name, last_name, mail, hashed_password)
	if err != nil {
		log.Fatal("Erreur lors de la récupération des utilisateurs : ", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Erreur lors de la récupération du nombre de lignes affectées : ", err)
		return
	}
	log.Printf("Nombre de lignes affectées : %d", rowsAffected)
	c.JSON(http.StatusOK, gin.H{
		"Message":                "Utilisateur créé avec succès",
		"Nombre de lignes créés": rowsAffected,
	})
}

func AdminModify(c *gin.Context) {
	if Id_user == "1" {
		var NewUser map[string]string
		id := c.Param("idd")
		if err := c.ShouldBindJSON(&NewUser); err != nil {
			c.JSON(400, gin.H{"Error": err.Error()})
			return
		}
		var first_name, last_name, mail, password string
		if val, ok := NewUser["first_name"]; ok {
			first_name = val
		}
		if val, ok := NewUser["last_name"]; ok {
			last_name = val
		}
		if val, ok := NewUser["mail"]; ok {
			mail = val
		}
		if val, ok := NewUser["password"]; ok {
			password = val
		}
		hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Erreur lors du hachage du mot de passe"})
			return
		}
		result, err := config.DB.ExecContext(c,
			"UPDATE users SET first_name = ?, last_name = ?, mail = ?, password = ? WHERE id = ?", first_name, last_name, mail, hashed_password, id)
		if err != nil {
			log.Fatal("Erreur lors de la récupération des utilisateurs : ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Erreur lors de la récupération du nombre de lignes affectées : ", err)
			return
		}
		log.Printf("Nombre de lignes affectées : %d", rowsAffected)
		c.JSON(http.StatusOK, gin.H{
			"Message":                   "Utilisateur modifiée avec succès",
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

func UserModify(c *gin.Context) {
	if Id_user != "" {
		var NewUser map[string]string
		if err := c.ShouldBindJSON(&NewUser); err != nil {
			c.JSON(400, gin.H{"Error": err.Error()})
			return
		}
		var first_name, last_name, mail, password string
		if val, ok := NewUser["first_name"]; ok {
			first_name = val
		}
		if val, ok := NewUser["last_name"]; ok {
			last_name = val
		}
		if val, ok := NewUser["mail"]; ok {
			mail = val
		}
		if val, ok := NewUser["password"]; ok {
			password = val
		}
		hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Erreur lors du hachage du mot de passe"})
			return
		}
		result, err := config.DB.ExecContext(c,
			"UPDATE users SET first_name = ?, last_name = ?, mail = ?, password = ? WHERE id = ?", first_name, last_name, mail, hashed_password, Id_user)
		if err != nil {
			log.Fatal("Erreur lors de la récupération des utilisateurs : ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Erreur lors de la récupération du nombre de lignes affectées : ", err)
			return
		}
		log.Printf("Nombre de lignes affectées : %d", rowsAffected)
		c.JSON(http.StatusOK, gin.H{
			"Message":                   "Utilisateur modifié avec succès",
			"Nombre de lignes modifiés": rowsAffected,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Veuillez vous connecter"})
	}
}

func AdminDelete(c *gin.Context) {
	if Id_user == "1" {
		id := c.Param("idd")
		result, err := config.DB.ExecContext(c,
			"DELETE FROM users WHERE id = ?", id)
		if err != nil {
			log.Fatal("Erreur lors de la récupération des utilisateurs : ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Erreur lors de la récupération du nombre de lignes affectées : ", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message":                    "Utilisateur supprimée avec succès",
			"nombre de lignes supprimés": rowsAffected,
		})
	} else if Id_user > "1" {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Permission refusée"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Veuillez vous connecter"})
	}
}

func UserDelete(c *gin.Context) {
	if Id_user != "" {
		result, err := config.DB.ExecContext(c,
			"DELETE FROM users WHERE id = ?", Id_user)
		if err != nil {
			log.Fatal("Erreur lors de la récupération des utilisateurs : ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Erreur lors de la récupération du nombre de lignes affectées : ", err)
			return
		}
		Id_user = ""
		c.JSON(http.StatusOK, gin.H{
			"Message":                    "Utilisateur supprimée avec succès",
			"Nombre de lignes supprimés": rowsAffected,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Veuillez vous connecter"})
	}
}

var Id_user string

func UserConnect(c *gin.Context) {
	var ConnectUser map[string]string
	var mail, password string
	if err := c.ShouldBindJSON(&ConnectUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Données de connexion invalides", "Details": err.Error()})
		return
	}
	if val, ok := ConnectUser["mail"]; ok {
		mail = val
	}
	if val, ok := ConnectUser["password"]; ok {
		password = val
	}
	var hashedPassword string
	err := config.DB.QueryRow("SELECT password FROM users WHERE mail = ?", mail).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Utilisateur non trouvé"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Erreur lors de la récupération des données utilisateur"})
		}
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Mot de passe incorrect"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Authentification réussie"})
	config.DB.QueryRow("SELECT id FROM users WHERE mail = ?", mail).Scan(&Id_user)
}
