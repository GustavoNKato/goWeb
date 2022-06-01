package controller

import (
	"github.com/GustavoNKato/goWeb/internal/products"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type request struct {
	ProductName string  `json:"product_name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Amount      int     `json:"amount"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "1234" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		product, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "1234" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		product, err := c.service.Store(req.ProductName, req.Color, req.Price, req.Amount)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

func (c *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "1234" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
			return
		}
		product, err := c.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "1234" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var request request
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		product, err := c.service.Update(id, request.ProductName, request.Color, request.Price, request.Amount)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, product)
	}
}
