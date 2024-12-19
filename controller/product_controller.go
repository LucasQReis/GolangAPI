package controller

import (
	"GolangAPI/useCase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	productUseCase useCase.ProductUseCase
}

func NewProductController(useCase useCase.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: useCase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}
