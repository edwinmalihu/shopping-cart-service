package controller

import (
	"net/http"
	"shopping-cart-service/repository"
	"shopping-cart-service/request"
	"shopping-cart-service/response"
	"shopping-cart-service/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController interface {
	AddCart(*gin.Context)
	ListCart(*gin.Context)
	DeleteCart(*gin.Context)
}

type cartController struct {
	cartRepo repository.CartRepo
}

// AddCart implements CartController.
func (c cartController) AddCart(ctx *gin.Context) {
	var req request.RequestAddCart
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := c.cartRepo.DetailProduct(req.ProductID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	isTrue := utils.ValidateStok(product.Stok, req.Quantity)
	if isTrue {
		ctx.JSON(http.StatusBadRequest, "Data yang dimasukan salah")
		return
	}

	data, err := c.cartRepo.AddCart(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.cartRepo.DetailCart(data.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := response.ResponseSuccess{
		ProductID: result.ProductID,
		Name:      result.Name,
		Price:     result.Price,
		Quantity:  req.Quantity,
		Msq:       "Data Berhasil di Tambahkan",
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteCart implements CartController.
func (c cartController) DeleteCart(ctx *gin.Context) {
	var req request.RequesCardById
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.Atoi(req.Id)

	data, err := c.cartRepo.DeleteCart(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)

}

// ListCart implements CartController.
func (c cartController) ListCart(ctx *gin.Context) {
	data, err := c.cartRepo.ListCart()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)

}

func NewCartController(repo repository.CartRepo) CartController {
	return cartController{
		cartRepo: repo,
	}
}
