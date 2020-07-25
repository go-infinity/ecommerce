package main

import "fmt"

func main() {

	// team := make([]string, 5) "Z"
	team1 := []string{"D", "S", "A", "V", "M"}
	team2 := append(team1[:2], team1[3:]...)

	// x, team1 := team1[len(team1)-1], team1[:len(team1)-1]
	// team2 := []string{"X", "Y", "Z"}

	// // Slice[start:end)
	// team3 := team1[:]
	// team := append(team1, team2...) //=>append(team1, "X,"Y","Z")
	// fmt.Printf("Popped Element is %s\n", x)
	fmt.Printf("Team %v\n", team2)
	// fmt.Printf("Team is %v,has length %d, has capactiy %d\n", team, len(team), cap(team))

	// productTVLG := product.Product{
	// 	ProductName:     "LG-TV1",
	// 	ProductCategory: product.ProductCategory{ParentCategoryName: "Electronics", SubCategory: "TV"},
	// 	ProductPrice:    10000,
	// 	ProductStock:    10,
	// 	LastAddedTime:   time.Now().Add(time.Hour * 10)}
	// productTVSAMSUNG := product.Product{
	// 	ProductName:     "SAMSUNG-TV1",
	// 	ProductCategory: product.ProductCategory{ParentCategoryName: "Electronics", SubCategory: "TV"},
	// 	ProductPrice:    10000,
	// 	ProductStock:    10,
	// 	LastAddedTime:   time.Now().Add(time.Hour * 10)}

	// products := []product.Product{productTVLG, productTVSAMSUNG}

	// var declare and intialize later
	// var user user.User
	// var declare in initzile in same statement
	// productTVLG := product.Product{
	// 	ProductName:     "LG-TV1",
	// 	ProductCategory: product.ProductCategory{ParentCategoryName: "Electronics", SubCategory: "TV"},
	// 	ProductPrice:    10000,
	// 	ProductStock:    10,
	// 	LastAddedTime:   time.Now().Add(time.Hour * 10)}
	// productTVSAMSUNG := product.Product{
	// 	ProductName:     "SAMSUNG-TV1",
	// 	ProductCategory: product.ProductCategory{ParentCategoryName: "Electronics", SubCategory: "TV"},
	// 	ProductPrice:    10000,
	// 	ProductStock:    10,
	// 	LastAddedTime:   time.Now().Add(time.Hour * 10)}
	// productFridgeLG := product.Product{
	// 	ProductName:     "LG-FRIDGE1",
	// 	ProductCategory: product.ProductCategory{ParentCategoryName: "Electronics", SubCategory: "Fridge"},
	// 	ProductPrice:    20000,
	// 	ProductStock:    5}
	// productFridgeSAMSUNG := product.Product{
	// 	ProductName:     "SAMSUNG-FRIDGE1",
	// 	ProductCategory: product.ProductCategory{ParentCategoryName: "Electronics", SubCategory: "Fridge"},
	// 	ProductPrice:    20000,
	// 	ProductStock:    5}
	// user := user.User{FirstName: "Sandeep",
	// 	LastName:        "Kumar",
	// 	DeliveryAddress: user.Address{City: "ModiNagar", State: user.UP, Country: user.INDIA},
	// 	IsActive:        true,
	// 	LastLoginTime:   time.Now()}

	// fmt.Printf("User is %v\n", user)

	//Data structures
	//Array
	// intArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Printf("the array is %v and length is %d", intArray, len(intArray))
	//Map
	// mapPopulation := map[string]int{"India": 25, "France": 5}
	// fmt.Printf("the Map is %v\n", mapPopulation)

	// products := [2]product.Product{productFridge, productTV}
	// products := map[string][]*product.Product{
	// 	"TV": []*product.Product{&productTVLG, &productTVSAMSUNG},
	// 	//        []string{"TV","VCR","DVD"}
	// 	"Fridge": []*product.Product{&productFridgeLG, &productFridgeSAMSUNG},
	// }
	// success := products["TV"].BuyProduct("TV")
	// success = products["TV"].BuyProduct("TV")
	// f.Printf("operation is %v\n", success)

	// success = products["Fridge"].BuyProduct("Fridge")
	// f.Printf("operation is %v\n", success)
	// for key, value := range products {
	// 	fmt.Printf("%s has value %v\n", key, *value)
	// }
}
