package models

type Purchase struct {
	Id                int    `json:"id"`
	User              int    `json:"user"`
	Quantity          int    `json:"quantity"`
	Total_price       int    `json:"total_price"`
	Payment_timestamp string `json:"payment_timestamp"`
	Book              int    `json:"book"`
}
