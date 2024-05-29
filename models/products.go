package models

type Product struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	PictureURL  string  `json:"picture"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}