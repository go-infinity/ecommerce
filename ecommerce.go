package main

import (
	"fmt"
	"time"

	book "github.com/go-infinity/ecommerce/Book"
)

var CompanyName string
var ExchangeRate int

func init() {
	CompanyName = "Go-Infinity"
}
func init() {
	ExchangeRate = 70
}

func ToUSD(amountINR float64) float64 {
	return amountINR / 70
}

func ToUSDAndPound(amountINR float64) (float64, float64) {
	return amountINR / 70, amountINR / 90
}

func ToUSDWithSymbol(amountINR float64) (amountUSD float64, symbol string) {
	return amountINR / 70, "$"
}

// SumInt takes any number of integer arguments and return sum of all integers
func SumInt(ints ...int) int {
	sum := 0
	for _, val := range ints {
		sum = sum + val
	}
	return sum
}

func main() {
	// var HelloFunc = func(audience string) string {
	// 	return "Hello " + audience
	// }
	// fmt.Printf("%s\n\n", HelloFunc("Go-Infinity"))
	// fmt.Printf("The name of the company is %s and has exchange rate of %d\n", CompanyName, ExchangeRate)
	// // team := make([]string, 5) "Z"
	// team1 := []string{"D", "S", "A", "V", "M"}
	// team2 := append(team1[:2], team1[3:]...)

	// // x, team1 := team1[len(team1)-1], team1[:len(team1)-1]
	// // team2 := []string{"X", "Y", "Z"}

	// // // Slice[start:end)
	// // team3 := team1[:]
	// // team := append(team1, team2...) //=>append(team1, "X,"Y","Z")
	// // fmt.Printf("Popped Element is %s\n", x)
	// fmt.Printf("Team %v\n", team2)
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
	// USDValue, USDSymbol := ToUSDWithSymbol(70)
	// fmt.Printf("Value of 70 INR is %.2f%s\n", USDValue, USDSymbol)

	// fmt.Printf("Sum of 1,2,3,4,5 is %d", SumInt(1, 2, 3, 4, 5))

	book := &book.Book{BookID: 1,
		BookTitle:   "Golang",
		BookAuthor:  "Vikash",
		BookSubject: "Programming"}

	fmt.Println(book.BookTitle)
	//Synchronous Way and Asynchornus
	go book.UpdateTitle("Go") //book - address - 0xc00004c040 // No need to handle return values
	go book.UpdateTitle("Go Programming")
	time.Sleep(time.Second * 1)
	fmt.Println(book.BookTitle) //book - address - 0xc00004c040
}
