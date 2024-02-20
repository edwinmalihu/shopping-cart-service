package response

type ResponseSuccess struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  uint    `json:"qty"`
	Msq       string  `json:"msg"`
}

type ResponseCart struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"total_amount"`
	Quantity  uint    `json:"qty"`
}
