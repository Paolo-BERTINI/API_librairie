package models

type Author struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Birth_date  string `json:"birth_date"`
	Description string `json:"description"`
}

type Author_list struct {
	Name        string `json:"name"`
	Birth_date  string `json:"birth_date"`
	Description string `json:"description"`
	Id          string `json:"id"`
}

type Author_name struct {
	Name string `json:"name"`
}
