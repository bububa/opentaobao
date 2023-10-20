package ascpchannel

import (
	"sync"
)

// TopChannelSkuPrice 结构体
type TopChannelSkuPrice struct {
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// sku价格
	SkuPrice *SkuPrice `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
}

var poolTopChannelSkuPrice = sync.Pool{
	New: func() any {
		return new(TopChannelSkuPrice)
	},
}

// GetTopChannelSkuPrice() 从对象池中获取TopChannelSkuPrice
func GetTopChannelSkuPrice() *TopChannelSkuPrice {
	return poolTopChannelSkuPrice.Get().(*TopChannelSkuPrice)
}

// ReleaseTopChannelSkuPrice 释放TopChannelSkuPrice
func ReleaseTopChannelSkuPrice(v *TopChannelSkuPrice) {
	v.SkuId = ""
	v.SkuPrice = nil
	poolTopChannelSkuPrice.Put(v)
}
