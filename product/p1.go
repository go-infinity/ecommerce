package product

import "time"

//CONSTANT
//ENUM
//CUSTOM TYPE
//VALIDATION

//Product is a type for Business products ABC
type Product struct {
	ProductName     string
	ProductCategory ProductCategory // ELECTRONICS,BEAUTY,GARMENT
	ProductPrice    float32
	ProductStock    int
	LastAddedTime   time.Time
}

// BuyProduct will update product stock
func (p *Product) BuyProduct(prodName string) bool {
	if p.ProductName == prodName {
		p.ProductStock = p.ProductStock - 1
		return true
	}

	return false
}
