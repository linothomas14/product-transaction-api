package main

import (
	// "github.com/linothomas14/product-transaction-api/config"
	// "github.com/linothomas14/product-transaction-api/controller"
	// "github.com/linothomas14/product-transaction-api/repository"
	// "github.com/linothomas14/product-transaction-api/service"

	"github.com/gin-gonic/gin"
	"github.com/linothomas14/product-transaction-api/config"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

// 	productRepository repository.ProductRepository = repository.NewProductRepository(db)
// 	transactionRepository  repository.TransactionRepository  = repository.NewTransactionRepository(db)

// 	productService service.ProductService = service.NewProductService(productRepository)
// 	transactionService  service.TransactionService  = service.NewTransactionService(transactionRepository, productRepository)

// 	productController controller.ProductController = controller.NewProductController(productService)
// 	transactionController  controller.TransactionController  = controller.NewTransactionController(transactionService)
)

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
}

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	// mhsRoutes := r.Group("products")
	// {
	// 	mhsRoutes.GET("/", productController.FindAll)
	// 	mhsRoutes.GET("/:id", productController.FindById)
	// 	mhsRoutes.POST("/", productController.Insert)
	// 	mhsRoutes.PUT("/:id", productController.Edit)
	// 	mhsRoutes.DELETE("/:id", productController.Delete)
	// }

	// transactionRoutes := r.Group("transactions")
	// {
	// 	transactionRoutes.GET("/", transactionController.FindAll)
	// 	transactionRoutes.GET("/:id", transactionController.FindById)
	// 	transactionRoutes.POST("/", transactionController.Insert)
	// 	transactionRoutes.PUT("/:id", productController.FindbyID)
	// 	transactionRoutes.DELETE("/:id", transactionController.Delete)

	// }
	r.GET("ping", PingHandler)
	r.Run()
}
