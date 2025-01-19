package models

type User struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Mail       string `json:"mail"`
	Password   string `json:"password"`
}

var Id_user string

type User_info struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Mail       string `json:"mail"`
}

type User_info_admin struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Mail       string `json:"mail"`
}
