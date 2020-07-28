package product

import (
	"time"
)

type Product struct {
	ProductName     string
	ProductCategory string
	ProductPrice    float32
	ProductStock    int32
	LastAdded       time.Time
}

// BuyProduct will update product stock
func (p *Product) BuyProduct(prodName string) bool {
	if p.ProductName == prodName {
		p.ProductStock = p.ProductStock - 1
		return true
	}
	return false
}
