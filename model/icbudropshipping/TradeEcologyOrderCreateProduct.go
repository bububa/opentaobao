package icbudropshipping

import (
	"sync"
)

// TradeEcologyOrderCreateProduct 结构体
type TradeEcologyOrderCreateProduct struct {
	// quantity
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// sku id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// unit price
	UnitPriceStr string `json:"unit_price_str,omitempty" xml:"unit_price_str,omitempty"`
	// product id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolTradeEcologyOrderCreateProduct = sync.Pool{
	New: func() any {
		return new(TradeEcologyOrderCreateProduct)
	},
}

// GetTradeEcologyOrderCreateProduct() 从对象池中获取TradeEcologyOrderCreateProduct
func GetTradeEcologyOrderCreateProduct() *TradeEcologyOrderCreateProduct {
	return poolTradeEcologyOrderCreateProduct.Get().(*TradeEcologyOrderCreateProduct)
}

// ReleaseTradeEcologyOrderCreateProduct 释放TradeEcologyOrderCreateProduct
func ReleaseTradeEcologyOrderCreateProduct(v *TradeEcologyOrderCreateProduct) {
	v.Quantity = ""
	v.SkuId = ""
	v.UnitPriceStr = ""
	v.ProductId = 0
	poolTradeEcologyOrderCreateProduct.Put(v)
}
