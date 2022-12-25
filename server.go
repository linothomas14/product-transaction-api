package main

import (
	"github.com/linothomas14/product-transaction-api/controller"
	"github.com/linothomas14/product-transaction-api/middleware"
	"github.com/linothomas14/product-transaction-api/repository"
	"github.com/linothomas14/product-transaction-api/service"

	"github.com/gin-gonic/gin"
	"github.com/linothomas14/product-transaction-api/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	userRepository repository.UserRepository = repository.NewUserRepository(db)
	// 	productRepository repository.ProductRepository = repository.NewProductRepository(db)
	// 	transactionRepository  repository.TransactionRepository  = repository.NewTransactionRepository(db)
	// jwtService  service.JWTService  = service.NewJWTService()
	authService service.AuthService = service.NewAuthService(userRepository)
	jwtService  service.AuthService = service.NewAuthService(userRepository)
	// 	productService service.ProductService = service.NewProductService(productRepository)
	// 	transactionService  service.TransactionService  = service.NewTransactionService(transactionRepository, productRepository)

	authController controller.AuthController = controller.NewAuthController(authService)

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

	r.POST("/login", authController.Login)
	r.POST("/register", authController.Register)

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", PingHandler)
		// userRoutes.PUT("/profile", userController.Update)
	}

	r.GET("ping", PingHandler)
	r.Run()
}
