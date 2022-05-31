package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {

	router := gin.Default()
	router.GET("/db", addData)
	router.GET("/products", productHandler)
	router.GET("/products/:id", getProductByIdHandler)
	router.POST("/products", createProduct)
	router.Run()

}

var lastID int
var products []product

type product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Inventory   int     `json:"inventory"`
	ProductCode string  `json:"product_code"`
	Publication bool    `json:"publication"`
	Data        string  `json:"data"`
}

func productHandler(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, products)
}

func getProductByIdHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("precisa ser um int - ", err)
	}

	for _, product := range products {
		if product.Id == id {
			ctx.IndentedJSON(http.StatusOK, product)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "produto não encontrado"})

}

func createProduct(ctx *gin.Context) {
	var newProduct product
	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	lastID++
	newProduct.Id = lastID

	products = append(products, newProduct)
	ctx.IndentedJSON(http.StatusCreated, newProduct)
}

func addData(ctx *gin.Context) {
	products = readJsonFile(products)
	lastID = 3
	ctx.IndentedJSON(http.StatusOK, products)
}

func readJsonFile(products []product) []product {
	jsonFile, err := ioutil.ReadFile("products.json")
	if err != nil {
		log.Println("erro ao ler o arquivo -", err)
	}

	if err := json.Unmarshal(jsonFile, &products); err != nil {
		log.Println("erro no momento do decode no arquivo json -", err)
	}
	return products
}
