package service

import (
	"github.com/linothomas14/product-transaction-api/model"
	"github.com/linothomas14/product-transaction-api/repository"
)

type ProductService interface {
	GetAllProduct() ([]model.Product, error)
	AddProduct(product model.Product) (model.Product, error)
	// GetProductByName(productName string) (model.Product, error)
	// GetProductById(id uint32) (model.Product, error)
	// UpdateProduct(product model.Product) (model.Product, error)
	// DeleteProduct(idProduct string) error
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRep repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRep,
	}
}

func (service *productService) GetAllProduct() ([]model.Product, error) {
	var products []model.Product

	products, err := service.productRepository.GetAllProduct(products)

	return products, err
}

func (service *productService) AddProduct(product model.Product) (model.Product, error) {

	product, err := service.productRepository.AddProduct(product)

	return product, err
}
