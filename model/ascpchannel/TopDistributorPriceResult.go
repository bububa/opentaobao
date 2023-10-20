package ascpchannel

import (
	"sync"
)

// TopDistributorPriceResult 结构体
type TopDistributorPriceResult struct {
	// sku价格详情
	SkuPrices []TopChannelSkuPrice `json:"sku_prices,omitempty" xml:"sku_prices>top_channel_sku_price,omitempty"`
	// 产品ID
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 产品价格详情
	ProductPrice *TopChannelPriceDetail `json:"product_price,omitempty" xml:"product_price,omitempty"`
}

var poolTopDistributorPriceResult = sync.Pool{
	New: func() any {
		return new(TopDistributorPriceResult)
	},
}

// GetTopDistributorPriceResult() 从对象池中获取TopDistributorPriceResult
func GetTopDistributorPriceResult() *TopDistributorPriceResult {
	return poolTopDistributorPriceResult.Get().(*TopDistributorPriceResult)
}

// ReleaseTopDistributorPriceResult 释放TopDistributorPriceResult
func ReleaseTopDistributorPriceResult(v *TopDistributorPriceResult) {
	v.SkuPrices = v.SkuPrices[:0]
	v.ProductId = ""
	v.ProductPrice = nil
	poolTopDistributorPriceResult.Put(v)
}
