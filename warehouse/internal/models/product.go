package models

type Product struct {
	ID      int     `json:"id"`
	Article string  `json:"article"`
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
}

type PostProduct struct {
	Article string  `json:"article"`
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
}
