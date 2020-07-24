package main

import (
	"fmt"
	"time"

	"github.com/go-infinity/ecommerce/order"
)

//func main() {
// var declare and intialize later
// var user user.User
// var declare in initzile in same statement
//product := product.Product{ProductName: "TV",
//	ProductCategory: "Electronics",
//	ProductPrice:    10000,
//	ProductStock:    10, LastAdded:time.Now().Add}
// user := user.User{FirstName: "Sandeep",
// 	LastName:        "Kumar",
// 	DeliveryAddress: user.Address{City: "ModiNagar", State: user.UP, Country: user.INDIA},
// 	IsActive:        true,
// 	LastLoginTime:   time.Now()}

// fmt.Printf("User is %v\n", user)
//fmt.Printf("product is %v\n", product)
//fmt.Printf("Last Added Time IS:", time.Now().Add(-24*time.Hour))
//success := product.BuyProduct("VCR")
//fmt.Printf("operation is %v\n", success)
//fmt.Printf("product is %v\n", product)
func main() {
	order := order.Order{OrderNumber: 1234,
		OrderBy:              "Azazul",
		OrderDate:            time.Time.Now(),
		OrderDeliveryAddress: order.Address{HouseNo: 12, StreeNo: 3, AreaVillage: "Bhualpur", City: "Chapra", State: "Bihar", Landmark: "Azad Chowk", PinCode: 841415},
		OrderDelivered:       true}

	fmt.Println("Order Delivery Details%v\n", order)
	//fmt.Println("Your order no is:%v\n", OrderNumber)
	//fmt.Println("Order By:%v\n", OrderBy)
	//fmt.Println("Order Date:%vn", OrderDate)
	//fmt.Println("Your order is successfully deliverd..%v\n", OrderDelivered)
	//fmt.Println("Delivery Address:%v\n", HouseNo, StreetNo, AreaVillage, City, State, Landmark, PinCode)
}
