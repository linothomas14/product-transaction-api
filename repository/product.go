package repository

import (
	"github.com/linothomas14/product-transaction-api/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct(product []model.Product) ([]model.Product, error)
	AddProduct(product model.Product) (model.Product, error)
	GetProductByName(productName string) (model.Product, error)
	GetProductById(id uint32) (model.Product, error)
	UpdateProduct(product model.Product) (model.Product, error)
	DeleteProduct(idProduct string) error
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) GetAllProduct(product []model.Product) ([]model.Product, error) {
	var products []model.Product

	err := db.connection.Find(&products).Error

	return products, err
}

func (db *productConnection) AddProduct(product model.Product) (model.Product, error) {

	err := db.connection.Create(&product).Error

	return product, err
}

func (db *productConnection) GetProductByName(productName string) (model.Product, error) {

	var product model.Product

	err := db.connection.Where("name", productName).First(&product).Error

	return product, err
}

func (db *productConnection) GetProductById(id uint32) (model.Product, error) {

	var product model.Product

	err := db.connection.Where("id", id).First(&product).Error

	return product, err
}

func (db *productConnection) UpdateProduct(product model.Product) (model.Product, error) {

	err := db.connection.Save(&product).Error

	return product, err
}

func (db *productConnection) DeleteProduct(idProduct string) error {

	var product model.Product

	err := db.connection.Where("id = ?", idProduct).Delete(&product).Error

	return err
}
