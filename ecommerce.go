package main

import (
	"fmt"

	"github.com/go-infinity/ecommerce/product"
	"github.com/go-infinity/ecommerce/user"
)

func main() {
	// var declare and intialize later
	var user user.User
	// var declare in initzile in same statement
	product := product.Product{ProductName: "TV", ProductCategory: "Electronics", ProductPrice: 10000}

	user.FirstName = "vikash"
	user.LastName = "Kumar"
	user.Role = "GUEST"
	user.EmailID = "vikas.kumar@gmail.com"
	user.DeliveryAddress = "PATNA"
	user.HomeAddress = "PATNA"

	fmt.Printf("User is %v\n", user)
	fmt.Printf("product is %v\n", product)
}
