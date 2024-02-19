package request

type RequestAddCart struct {
	CustomerID uint `json:"customer_id"`
	ProductID  uint `json:"product_id"`
	Quantity   uint `json:"qty"`
}

type RequesCardById struct {
	Id string `form:"cart_id"`
}
