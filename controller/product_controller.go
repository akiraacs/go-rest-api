package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			ID:    1,
			Name:  "iPhone 13",
			Price: 999.99,
		},
		{
			ID:    2,
			Name:  "iPhone 13 Pro",
			Price: 1999.99,
		},
		{
			ID:    3,
			Name:  "iPhone 13 Pro Max",
			Price: 2999.99,
		},
	}

	ctx.JSON(http.StatusOK, products)
}