package fenxiao

import (
	"sync"
)

// TopProductPriceResult 结构体
type TopProductPriceResult struct {
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 币种
	CurrencyType string `json:"currency_type,omitempty" xml:"currency_type,omitempty"`
	// 产品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// SKU ID
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 渠道编码
	ChannelCode int64 `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 价格类型，区域价、指导价
	PriceType int64 `json:"price_type,omitempty" xml:"price_type,omitempty"`
}

var poolTopProductPriceResult = sync.Pool{
	New: func() any {
		return new(TopProductPriceResult)
	},
}

// GetTopProductPriceResult() 从对象池中获取TopProductPriceResult
func GetTopProductPriceResult() *TopProductPriceResult {
	return poolTopProductPriceResult.Get().(*TopProductPriceResult)
}

// ReleaseTopProductPriceResult 释放TopProductPriceResult
func ReleaseTopProductPriceResult(v *TopProductPriceResult) {
	v.Price = ""
	v.CurrencyType = ""
	v.ProductId = 0
	v.SkuId = 0
	v.ChannelCode = 0
	v.PriceType = 0
	poolTopProductPriceResult.Put(v)
}
