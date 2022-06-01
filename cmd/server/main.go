package main

import (
	"github.com/GustavoNKato/goWeb/cmd/server/controller"
	"github.com/GustavoNKato/goWeb/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	product := controller.NewProduct(service)

	router := gin.Default()
	productRouter := router.Group("/products")
	productRouter.POST("/", product.Store())
	productRouter.GET("/", product.GetAll())
	productRouter.GET("/:id", product.GetById())
	router.Run()
}
