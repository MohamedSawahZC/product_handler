package models

import (
	"crud/src/entities"
	"database/sql"
)

type ProductModel struct {
	Db *sql.DB
}

/*
============== Search by name ==============
 @desc Method to find all products
 @route /api/product/all
 @access Public
======================================
*/
func (productModel ProductModel) Search(keyword string) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product where name like ?", "%"+keyword+"%")

	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}

}

/*
============== Find All ==============
 @desc Method to find all products
 @route /api/product/all
 @access Public
======================================
*/

func (productModel ProductModel) FindAll() (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product")

	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}

}
