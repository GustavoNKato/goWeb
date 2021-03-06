package products

import (
	"fmt"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Amount      int     `json:"amount"`
}

var products []Product

var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Store(id int, productName, color string, price float64, amount int) (Product, error)
	LastID() (int, error)
	Update(id int, productName, color string, price float64, amount int) (Product, error)
}

type repository struct{}

func (r *repository) GetById(id int) (Product, error) {
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return Product{}, fmt.Errorf("produto nao encontrado")
}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (r *repository) Store(id int, productName, color string, price float64, amount int) (Product, error) {
	product := Product{id, productName, color, price, amount}
	products = append(products, product)
	lastID = product.Id
	return product, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Update(id int, productName, color string, price float64, amount int) (Product, error) {
	product := Product{ProductName: productName, Color: color, Price: price, Amount: amount}
	updated := false
	for i := range products {
		if products[i].Id == id {
			product.Id = id
			products[i] = product
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return product, nil
}

func NewRepository() Repository {
	return &repository{}
}
