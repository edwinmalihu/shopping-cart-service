package main

import (
	"shopping-cart-service/model"
	"shopping-cart-service/route"
)

func main() {

	db, _ := model.DBConnection()
	route.SetupRoutes(db)

}
