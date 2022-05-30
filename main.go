package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	router := gin.Default()
	router.GET("/products", productHandler)
	router.GET("/products/:id", getProductByIdHandler)
	router.Run()

}

type product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Inventory   int     `json:"inventory"`
	Code        string  `json:"code"`
	Publication bool    `json:"publication"`
	Data        string  `json:"data"`
}

func productHandler(ctx *gin.Context) {
	var products []product
	products = readJsonFile(products)
	ctx.IndentedJSON(http.StatusOK, products)
}

func getProductByIdHandler(ctx *gin.Context) {
	var products []product
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("precisa ser um int - ", err)
	}

	products = readJsonFile(products)

	for _, product := range products {
		if product.Id == id {
			ctx.IndentedJSON(http.StatusOK, product)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "produto não encontrado"})

}

func readJsonFile(products []product) []product {
	jsonFile, err := os.ReadFile("products.json")
	if err != nil {
		log.Println("erro ao ler o arquivo -", err)
	}

	if err := json.Unmarshal(jsonFile, &products); err != nil {
		log.Println("erro no momento do decode no arquivo json -", err)
	}
	return products
}
