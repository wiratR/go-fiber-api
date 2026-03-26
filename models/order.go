package models

type Order struct {
	ID     int    `json:"id"`
	Item   string `json:"item"`
	Amount int    `json:"amount"`
	Status string `json:"status"`
}
