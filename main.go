package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {

	router := gin.Default()
	router.GET("/products", productHandler)
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

	jsonFile, err := os.ReadFile("products.json")
	if err != nil {
		log.Println("Erro ao ler o arquivo -", err)

	}

	if err := json.Unmarshal(jsonFile, &products); err != nil {
		log.Println("Erro no momento do decode no arquivo json -", err)
	}

	ctx.IndentedJSON(http.StatusOK, products)
}
