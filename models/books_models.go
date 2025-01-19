package models

type Book struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	Author           int    `json:"authors"`
	Stock            int    `json:"stock"`
	Publication_date string `json:"publication_date"`
	Price            int    `json:"price"`
	Summary          string `json:"summary"`
}

type BookResponseAdmin struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  string `json:"stock"`
}

type BookResponse struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Id     string `json:"id"`
}

type BookResponses struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Summary  string `json:"summary"`
	Price    int    `json:"price"`
	Authorid string `json:"id"`
}

type BookResponsesAdmin struct {
	Title   string `json:"title"`
	Stock   string `json:"stock"`
	Author  string `json:"author"`
	Summary string `json:"summary"`
	Price   int    `json:"price"`
}
