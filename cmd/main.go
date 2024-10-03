package main

import (
	"go-rest-api/controller"
	"go-rest-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

    ProductUseCase := usecase.NewProductUseCase()

    //Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

    server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
