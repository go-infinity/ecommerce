package main

import (
	"errors"
	"fmt"
)

type customerror struct {
	kind      string
	op        string
	errortype string
	code      int
	message   string
}

func (ce customerror) Error() string {
	return "Kind:" + ce.kind + "Message:" + ce.message
}

var (
	ProductInventoryNotAccessible = customerror{kind: "Product", op: "Buy", code: 3, message: "Product Inventory is not accessible", errortype: "Temporary"}
	ProductNotFound               = customerror{kind: "Product", op: "Buy", code: 1, message: "Product does not exist as per provided ID", errortype: "Permanent"}
	ProductDamaged                = customerror{kind: "Product", op: "Sell", code: 2, message: "Product is damaged as per the photos", errortype: "Permanent"}
	// ProductNotFound = errors.New("Product does not exist as per provided ID")
	// ProductDamaged  = errors.New("Product is damaged as per the photos")
)

type productInterface interface {
	BuyProduct(productID int) error
	SellProduct(productID int) error
}

// Amazon
type productImpl1 struct {
	productID     int
	productName   string
	productAmount float32
}

func (p *productImpl1) BuyProduct(productID int) error {
	return ProductNotFound
}

func (p *productImpl1) SellProduct(productID int) error {
	return ProductDamaged
}

func (p *productImpl1) ProductAmount(productID int) float32 {
	return 0
}

// Flipkart
type productImpl2 struct {
	productID       int
	productName     string
	productAmount   float32
	productQuantity int
	offercode       int
}

func (p *productImpl2) BuyProduct(productID int) error {

	return ProductNotFound
}

func (p *productImpl2) SellProduct(productID int) error {
	return ProductDamaged
}

func (p *productImpl2) ProductAmount(productID int) float32 {
	return 0
}
func BuyProduct(pID int, p productInterface) error {
	return p.BuyProduct(pID)
}

func main() {

	p1 := productImpl1{productID: 1, productName: "TV", productAmount: 100}
	p2 := productImpl2{productID: 2, productName: "VCR", productAmount: 200, offercode: 123, productQuantity: 10}
	fmt.Printf("The result is %v\n", BuyProduct(p1.productID, &p1))
	fmt.Printf("The result is %v\n", BuyProduct(p2.productID, &p2))

}

func addtworealnumber(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		// return 0, fmt.Errorf("One of the value is less than 0")
		return 0, errors.New("One of the value is less than 0")
	} else {
		return a + b, nil
	}
}

// main - > A -> B(error) -> C(error)

//  try {
//
//   os.Openfile
//
//}
//  catch {}
//
//
//
