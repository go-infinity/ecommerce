package main

import (
	"fmt"

	"github.com/go-infinity/ecommerce/product"
)

func main() {
	// var declare and intialize later
	// var user user.User
	// var declare in initzile in same statement
	product := product.Product{ProductName: "TV", ProductCategory: "Electronics", ProductPrice: 10000, ProductStock: 10}

	// user := user.User{FirstName: "Sandeep",
	// 	LastName:        "Kumar",
	// 	DeliveryAddress: user.Address{City: "ModiNagar", State: user.UP, Country: user.INDIA},
	// 	IsActive:        true,
	// 	LastLoginTime:   time.Now()}

	// fmt.Printf("User is %v\n", user)
	fmt.Printf("product is %v\n", product)
	success := product.BuyProduct("VCR")
	fmt.Printf("operation is %v\n", success)
	fmt.Printf("product is %v\n", product)
}
