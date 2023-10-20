package ascpchannel

import (
	"sync"
)

// SkuPrice 结构体
type SkuPrice struct {
	// 扩展价格
	ExtendPrice string `json:"extend_price,omitempty" xml:"extend_price,omitempty"`
	// 价格值
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 币种
	CurrencyPriceValue string `json:"currency_price_value,omitempty" xml:"currency_price_value,omitempty"`
}

var poolSkuPrice = sync.Pool{
	New: func() any {
		return new(SkuPrice)
	},
}

// GetSkuPrice() 从对象池中获取SkuPrice
func GetSkuPrice() *SkuPrice {
	return poolSkuPrice.Get().(*SkuPrice)
}

// ReleaseSkuPrice 释放SkuPrice
func ReleaseSkuPrice(v *SkuPrice) {
	v.ExtendPrice = ""
	v.Price = ""
	v.CurrencyPriceValue = ""
	poolSkuPrice.Put(v)
}
