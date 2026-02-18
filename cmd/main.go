package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// inicializa a conexão com o banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// camada de repository
	productRepository := repository.NewRepositoryProduct(dbConnection)

	// camada usecase
	ProductUseCase := usecase.NewProductUseCase(productRepository)

	// camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.DELETE("/product/:productId", ProductController.DeleteProductById)
	server.PUT("/product/:productId", ProductController.UpdateProduct)

	server.Run(":8000")
}
