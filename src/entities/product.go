package entities

import "fmt"

type Product struct {
	Id       int64
	Name     string
	Price    float64
	Quantity int64
}

func (product Product) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\nprice: %0.1fquantity: %d", product.Id, product.Name, product.Price, product.Quantity)
}
