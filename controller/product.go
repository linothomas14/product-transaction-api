package controller

import (
	"net/http"

	"github.com/linothomas14/product-transaction-api/helper"
	"github.com/linothomas14/product-transaction-api/model"
	"github.com/linothomas14/product-transaction-api/service"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetAllProduct(context *gin.Context)
	AddProduct(context *gin.Context)
}

type productController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (c *productController) GetAllProduct(ctx *gin.Context) {

	products, err := c.productService.GetAllProduct()

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	response := helper.BuildResponse("OK", products)
	ctx.JSON(http.StatusOK, response)

	return
}

func (c *productController) AddProduct(ctx *gin.Context) {

	var product model.Product

	err := ctx.ShouldBind(&product)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err = c.productService.AddProduct(product)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse("OK", product)
	ctx.JSON(http.StatusOK, response)
	return
}
