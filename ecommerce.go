package main

import (
	"fmt"
	"time"

	"github.com/go-infinity/ecommerce/order"
	"github.com/go-infinity/ecommerce/product"
)

func main() {
	product := product.Product{ProductName: "TV", ProductPrice: 764, ProductCategory: "Electronic", ProductStock: 5}
	fmt.Printf("product is %v\n", product)
	success := product.BuyProduct("TV")
	fmt.Printf("operation is %v\n", success)
	fmt.Printf("product is %v\n", product)
	order := order.Order{OrderNumber: 1234,
		OrderBy:         "Azazul",
		OrderDate:       time.Now(),
		DeliveryAddress: order.Address{HouseNo: 12, StreetNo: 3, AreaVillage: "Bhualpur", City: "Chapra", State: "Bihar", Landmark: "Azad Chowk", PinCode: 841415},
		OrderDelivered:  true}

	fmt.Printf("Order Delivery Details%v", order)
	//fmt.Println("Your order no is:%v\n", OrderNumber)
	//fmt.Println("Order By:%v\n", OrderBy)
	//fmt.Println("Order Date:%vn", OrderDate)
	//fmt.Println("Your order is successfully deliverd..%v\n", OrderDelivered)
	//fmt.Println("Delivery Address:%v\n", HouseNo, StreetNo, AreaVillage, City, State, Landmark, PinCode)
}
